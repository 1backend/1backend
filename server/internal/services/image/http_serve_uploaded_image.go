/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package imageservice

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	stdimage "image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/gorilla/mux"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/chai2010/webp"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"

	image "github.com/1backend/1backend/server/internal/services/image/types"
)

type ErrorResponse image.ErrorResponse

type imgResult struct {
	Data        []byte
	ContentType string
}

// @ID serveUploadedImage
// @Summary Serve Uploaded Image
// @Description Retrieves and serves a previously uploaded image file using its File ID.
// @Tags Image Svc
// @Accept json
// @Produce application/octet-stream
// @Param fileId path string true "FileID uniquely identifies the file itself (not an ID, which represents a specific replica)"
// @Param width query int false "Optional width to resize the image to"
// @Param height query int false "Optional height to resize the image to"
// @Success 200 {file} binary "File served successfully"
// @Failure 400 {object} image.ErrorResponse "Missing File ID"
// @Failure 404 {object} image.ErrorResponse "File Not Found"
// @Failure 500 {object} image.ErrorResponse "Internal Server Error"
// @Router /image-svc/serve/upload/{fileId} [get]
func (cs *ImageService) ServeUploadedImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileId := vars["fileId"]
	if fileId == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Missing file ID")
		return
	}

	// 1. Parse Params
	width, _ := strconv.Atoi(strings.TrimSuffix(r.URL.Query().Get("width"), "px"))
	height, _ := strconv.Atoi(strings.TrimSuffix(r.URL.Query().Get("height"), "px"))
	quality, _ := strconv.Atoi(r.URL.Query().Get("quality"))
	if quality <= 0 {
		quality = 85
	}

	// 2. Initial cache check (L1/L2)
	// We use a stable hash that doesn't rely on the metadata being found yet
	cacheKeyData := fmt.Sprintf("%s-%d-%d-%d", fileId, width, height, quality)
	h := sha1.New()
	h.Write([]byte(cacheKeyData))
	hash := hex.EncodeToString(h.Sum(nil))

	contentType, _ := cs.metaCache.Get(fileId)

	var rsp *os.File

	if contentType == "" {
		var (
			hrsp *http.Response
			err  error
		)

		rsp, hrsp, err = cs.options.ClientFactory.Client(client.WithToken(cs.token)).
			FileSvcAPI.ServeUpload(context.Background(), fileId).Execute()
		if err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		contentType = hrsp.Header["Content-Type"][0]
		cs.metaCache.Add(fileId, contentType)
	}

	// Check RAM
	if data, ok := cs.imageDataCache.Get(hash); ok {
		contentType, _ := cs.metaCache.Get(fileId)
		if contentType == "" {
			contentType = "image/png"
		}
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		w.Write(data)
		return
	}

	// Check Disk
	cachePath := filepath.Join(cs.imageCacheFolder, hash)
	if data, err := os.ReadFile(cachePath); err == nil {
		if len(data) < 100*1024 {
			cs.imageDataCache.Add(hash, data)
		}
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		w.Write(data)
		return
	}

	// 3. Singleflight - Only one worker processes the image
	val, err, _ := cs.sf.Do(hash, func() (interface{}, error) {
		// Double check disk
		if data, err := os.ReadFile(cachePath); err == nil {
			return &imgResult{Data: data, ContentType: contentType}, nil
		}

		// Fetch original
		if rsp == nil {
			var (
				hrsp *http.Response
				err  error
			)

			rsp, hrsp, err = cs.options.ClientFactory.Client(client.WithToken(cs.token)).
				FileSvcAPI.ServeUpload(context.Background(), fileId).Execute()
			if err != nil {
				return nil, err
			}

			contentType = hrsp.Header["Content-Type"][0]
			cs.metaCache.Add(fileId, contentType)
		}

		// --- YOUR ORIGINAL DECODE BLOCK RESTORED ---
		var img stdimage.Image
		var decodeErr error
		switch contentType {
		case "image/png":
			img, decodeErr = png.Decode(rsp)
		case "image/jpeg", "image/jpg":
			img, decodeErr = jpeg.Decode(rsp)
		case "image/webp":
			img, decodeErr = webp.Decode(rsp)
		default:
			img, _, decodeErr = stdimage.Decode(rsp)
		}
		if decodeErr != nil {
			return nil, decodeErr
		}

		// 4. Resize
		if width > 0 || height > 0 {
			img = cs.smartResize(img, width, height, transform.Lanczos)
		}

		// 5. Encode
		buf := new(bytes.Buffer)
		var encodeErr error
		switch contentType {
		case "image/jpeg", "image/jpg":
			encodeErr = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
		case "image/webp":
			encodeErr = webp.Encode(buf, img, &webp.Options{Quality: float32(quality)})
		default:
			encodeErr = png.Encode(buf, img)
		}
		if encodeErr != nil {
			return nil, encodeErr
		}

		finalData := buf.Bytes()
		_ = os.WriteFile(cachePath, finalData, 0644)
		cs.metaCache.Add(fileId, contentType) // Store for L1/L2 hits

		if len(finalData) < 100*1024 {
			cs.imageDataCache.Add(hash, finalData)
		}

		return &imgResult{Data: finalData, ContentType: contentType}, nil
	})

	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	// 6. Serve the result to everyone (Winner and Waiters)
	res := val.(*imgResult)
	w.Header().Set("Content-Type", res.ContentType)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Write(res.Data)
}

// Helper to handle aspect ratio and prevent enlargement
func (cs *ImageService) smartResize(img stdimage.Image, targetW, targetH int, filter transform.ResampleFilter) stdimage.Image {
	bounds := img.Bounds()
	origW, origH := bounds.Dx(), bounds.Dy()

	// 1. Calculate the scaling factor
	ratioW := float64(targetW) / float64(origW)
	ratioH := float64(targetH) / float64(origH)

	var finalW, finalH int

	if targetW > 0 && targetH > 0 {
		// Use the smaller ratio to ensure the image fits inside the box
		ratio := ratioW
		if ratioH < ratioW {
			ratio = ratioH
		}
		finalW = int(float64(origW) * ratio)
		finalH = int(float64(origH) * ratio)
	} else if targetW > 0 {
		finalW = targetW
		finalH = int(float64(targetW) * float64(origH) / float64(origW))
	} else if targetH > 0 {
		finalH = targetH
		finalW = int(float64(targetH) * float64(origW) / float64(origH))
	} else {
		return img
	}

	// 2. Prevent enlargement (Don't upscale if target is bigger than source)
	if finalW >= origW || finalH >= origH {
		return img
	}

	return transform.Resize(img, finalW, finalH, filter)
}
