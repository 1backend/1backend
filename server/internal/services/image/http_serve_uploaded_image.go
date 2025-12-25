/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package imageservice

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	stdimage "image"
	"image/gif"
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

	"github.com/gen2brain/avif"

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

	// 1. Parse Params (Optimized: No need to trim "px" manually if using Atoi generally)
	width, _ := strconv.Atoi(strings.TrimSuffix(r.URL.Query().Get("width"), "px"))
	height, _ := strconv.Atoi(strings.TrimSuffix(r.URL.Query().Get("height"), "px"))
	quality, _ := strconv.Atoi(r.URL.Query().Get("quality"))
	if quality <= 0 {
		quality = 85
	}

	resampleFilter := transform.Lanczos

	// 2. Metadata Lookup (Avoid unnecessary Network calls)
	contentType, metaFound := cs.metaCache.Get(fileId)

	// 3. Disk Cache Check
	// We generate the hash without the content-type if not found,
	// but include it if we have it to ensure unique variants.
	cacheKeyData := fmt.Sprintf("%s-%d-%d-%d-%s", fileId, width, height, quality, contentType)
	h := sha1.New()
	h.Write([]byte(cacheKeyData))
	hash := hex.EncodeToString(h.Sum(nil))
	cachePath := filepath.Join(cs.imageCacheFolder, hash)

	if f, err := os.Open(cachePath); err == nil {
		defer f.Close()
		if metaFound {
			w.Header().Set("Content-Type", contentType)
		} else {
			// Sniff or fallback if meta was lost from RAM but file exists on disk
			w.Header().Set("Content-Type", "image/png")
		}
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		io.Copy(w, f)
		return
	}

	// 4. Singleflight: Prevent concurrent processing of the same image variant
	_, err, _ := cs.sf.Do(hash, func() (interface{}, error) {
		// Double check disk inside singleflight to catch the winner's work
		if _, err := os.Stat(cachePath); err == nil {
			return nil, nil // Someone else just finished it
		}

		// 5. CACHE MISS: Fetch from File Service
		rsp, hrsp, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
			FileSvcAPI.ServeUpload(context.Background(), fileId).
			Execute()
		if err != nil {
			return nil, err
		}

		// Update meta cache immediately
		fetchedType := hrsp.Header.Get("Content-Type")
		if fetchedType == "" {
			fetchedType = "image/png"
		}
		cs.metaCache.Add(fileId, fetchedType)
		contentType = fetchedType

		// 6. Decode
		var img stdimage.Image
		switch contentType {
		case "image/png":
			img, err = png.Decode(rsp)
		case "image/jpeg", "image/jpg":
			img, err = jpeg.Decode(rsp)
		case "image/gif":
			img, err = gif.Decode(rsp)
		case "image/webp":
			img, err = webp.Decode(rsp)
		case "image/avif":
			img, err = avif.Decode(rsp)
		default:
			img, _, err = stdimage.Decode(rsp)
		}
		if err != nil {
			return nil, err
		}

		// 7. Resize (only if necessary)
		if width > 0 || height > 0 {
			img = cs.smartResize(img, width, height, resampleFilter)
		}

		// 8. Save and Stream
		outFile, err := os.Create(cachePath)
		if err != nil {
			return nil, err
		}
		defer outFile.Close()

		// Set response headers
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

		// Use MultiWriter to save to disk and send to user simultaneously
		mw := io.MultiWriter(w, outFile)
		switch contentType {
		case "image/jpeg", "image/jpg":
			err = jpeg.Encode(mw, img, &jpeg.Options{Quality: quality})
		case "image/webp":
			err = webp.Encode(mw, img, &webp.Options{Quality: float32(quality)})
		case "image/gif":
			err = gif.Encode(mw, img, nil)
		default:
			err = png.Encode(mw, img)
		}
		return nil, err
	})

	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
	}
}

// Helper to handle aspect ratio and prevent enlargement
func (cs *ImageService) smartResize(img stdimage.Image, w, h int, filter transform.ResampleFilter) stdimage.Image {
	bounds := img.Bounds()
	origW, origH := bounds.Dx(), bounds.Dy()

	if w <= 0 {
		w = int(float64(h) * float64(origW) / float64(origH))
	}
	if h <= 0 {
		h = int(float64(w) * float64(origH) / float64(origW))
	}

	// Prevent enlargement
	if w >= origW || h >= origH {
		return img
	}

	return transform.Resize(img, w, h, filter)
}
