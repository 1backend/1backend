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
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/gorilla/mux"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"

	"github.com/chai2010/webp"

	image "github.com/1backend/1backend/server/internal/services/image/types"
)

type ErrorResponse image.ErrorResponse

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

	contentType, metaFound := cs.metaCache.Get(fileId)

	// 2. Generate Hash
	cacheKeyData := fmt.Sprintf("%s-%d-%d-%d-%s", fileId, width, height, quality, contentType)
	h := sha1.New()
	h.Write([]byte(cacheKeyData))
	hash := hex.EncodeToString(h.Sum(nil))

	// 3. L1 CACHE CHECK (Memory)
	if data, ok := cs.imageDataCache.Get(hash); ok {
		w.Header().Set("Content-Type", contentType) // contentType might be empty if !metaFound
		if !metaFound {
			w.Header().Set("Content-Type", "image/png")
		}
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		w.Write(data)
		return
	}

	// 4. L2 CACHE CHECK (Disk)
	cachePath := filepath.Join(cs.imageCacheFolder, hash)
	if f, err := os.Open(cachePath); err == nil {
		defer f.Close()

		// Optimization: Read into buffer to serve and potentially populate L1
		data, err := io.ReadAll(f)
		if err == nil {
			// Populate L1 if small (e.g., < 100KB)
			if len(data) < 100*1024 {
				cs.imageDataCache.Add(hash, data)
			}
			w.Header().Set("Content-Type", contentType)
			if !metaFound {
				w.Header().Set("Content-Type", "image/png")
			}
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			w.Write(data)
			return
		}
	}

	// 5. CACHE MISS (Singleflight)
	val, err, _ := cs.sf.Do(hash, func() (interface{}, error) {
		// Re-check disk in case someone else just finished it
		if data, err := os.ReadFile(cachePath); err == nil {
			return data, nil
		}

		rsp, hrsp, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
			FileSvcAPI.ServeUpload(context.Background(), fileId).Execute()
		if err != nil {
			return nil, err
		}

		fetchedType := hrsp.Header.Get("Content-Type")
		if fetchedType == "" {
			fetchedType = "image/png"
		}
		cs.metaCache.Add(fileId, fetchedType)
		contentType = fetchedType

		// 6. Decode & Process
		var img stdimage.Image
		switch contentType {
		case "image/png":
			img, err = png.Decode(rsp)
		case "image/jpeg", "image/jpg":
			img, err = jpeg.Decode(rsp)
		case "image/webp":
			img, err = webp.Decode(rsp)
		default:
			img, _, err = stdimage.Decode(rsp)
		}
		if err != nil {
			return nil, err
		}

		if width > 0 || height > 0 {
			img = cs.smartResize(img, width, height, transform.Lanczos)
		}

		// 7. Encode to Buffer (Capture bytes for L1 and Disk)
		buf := new(bytes.Buffer)
		switch contentType {
		case "image/jpeg", "image/jpg":
			jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
		case "image/webp":
			webp.Encode(buf, img, &webp.Options{Quality: float32(quality)})
		default:
			png.Encode(buf, img)
		}

		finalData := buf.Bytes()

		// 8. Save to Disk (Async-ish or via WriteFile)
		_ = os.WriteFile(cachePath, finalData, 0644)

		// 9. Populate L1 if small
		if len(finalData) < 100*1024 {
			cs.imageDataCache.Add(hash, finalData)
		}

		return finalData, nil
	})

	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	// --- CRITICAL FIX START ---
	// If we are here, 'val' contains the []byte from a fresh process or a concurrent winner
	data, ok := val.([]byte)
	if !ok || len(data) == 0 {
		endpoint.WriteErr(w, http.StatusInternalServerError, fmt.Errorf("failed to retrieve image data"))
		return
	}

	// Set headers again because the waiting callers didn't run the switch/hrsp logic
	w.Header().Set("Content-Type", contentType)
	if contentType == "" {
		w.Header().Set("Content-Type", "image/png")
	}
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	w.Write(data)
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
