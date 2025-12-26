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
	"image/gif"
	"image/jpeg"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/gen2brain/avif"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	"github.com/chai2010/webp"

	image "github.com/1backend/1backend/server/internal/services/image/types"
)

type ErrorResponse image.ErrorResponse

type imgResult struct {
	Data        []byte
	ContentType string
}

const memCacheLimit = 250 * 1024

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

	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	qualityStr := r.URL.Query().Get("quality")

	width := 0
	height := 0
	quality := 85

	var err error
	if widthStr != "" {
		widthStr = strings.TrimSuffix(widthStr, "px")
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			endpoint.WriteErr(w, http.StatusBadRequest, errors.New("invalid width"))
			return
		}
	}
	if heightStr != "" {
		heightStr = strings.TrimSuffix(heightStr, "px")
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			endpoint.WriteErr(w, http.StatusBadRequest, errors.New("invalid height"))
			return
		}
	}
	if qualityStr != "" {
		quality, err = strconv.Atoi(qualityStr)
		if err != nil {
			endpoint.WriteErr(w, http.StatusBadRequest, errors.New("invalid quality"))
			return
		}
	}

	contentType, _ := cs.metaCache.Get(fileId)
	cacheKeyData := fmt.Sprintf("%s-%d-%d-%d", fileId, width, height, quality)
	h := sha1.New()
	h.Write([]byte(cacheKeyData))
	hash := hex.EncodeToString(h.Sum(nil))

	var rsp *os.File

	if contentType == "" {
		var (
			hrsp *http.Response
			err  error
		)

		_, err, _ = cs.sf.Do(hash, func() (interface{}, error) {
			rsp, hrsp, err = cs.options.ClientFactory.Client(client.WithToken(cs.token)).
				FileSvcAPI.ServeUpload(context.Background(), fileId).Execute()
			if err != nil {
				endpoint.WriteErr(w, http.StatusInternalServerError, err)
				return nil, errors.Wrap(err, "error calling serve upload to get content type")
			}

			contentType = hrsp.Header["Content-Type"][0]
			cs.metaCache.Add(fileId, contentType)

			return nil, nil
		})

		if err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		defer rsp.Close()
	}

	switch contentType {
	case "image/jpeg", "image/jpg":
		w.Header().Set("Content-Type", "image/jpeg")
	case "image/gif":
		w.Header().Set("Content-Type", "image/gif")
	case "image/webp":
		w.Header().Set("Content-Type", "image/webp")
	case "image/avif":
		w.Header().Set("Content-Type", "image/avif")
	default:
		w.Header().Set("Content-Type", "image/png")
	}

	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

	// Check RAM
	if data, ok := cs.imageDataCache.Get(hash); ok {
		w.Write(data)
		return
	}

	// Check disk
	cachePath := cs.getCachePath(hash)
	if data, err := os.ReadFile(cachePath); err == nil {
		if len(data) < memCacheLimit {
			cs.imageDataCache.Add(hash, data)
		}

		w.Write(data)
		return
	}

	val, err, _ := cs.sf.Do(hash, func() (interface{}, error) {
		logger.Info("getting from file service", hash)

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
			defer rsp.Close()

			cs.metaCache.Add(fileId, hrsp.Header.Get("Content-Type"))
		}

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
		case "image/tiff":
			img, err = tiff.Decode(rsp)
		case "image/bmp":
			img, err = bmp.Decode(rsp)
		case "image/avif":
			img, err = avif.Decode(rsp)
		default:
			// fall back to generic
			img, _, err = stdimage.Decode(rsp)
		}
		if err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return nil, errors.Wrap(err, "decode err")
		}

		if width > 0 && height > 0 {
			bounds := img.Bounds()
			origWidth := bounds.Dx()
			origHeight := bounds.Dy()

			// Compute the target width/height while maintaining aspect ratio
			var targetWidth, targetHeight int

			if width > 0 && height > 0 {
				aspectRatio := float64(origWidth) / float64(origHeight)
				if float64(width)/aspectRatio <= float64(height) {
					targetWidth = width
					targetHeight = int(float64(width) / aspectRatio)
				} else {
					targetHeight = height
					targetWidth = int(float64(height) * aspectRatio)
				}
			} else if width > 0 {
				targetWidth = width
				targetHeight = int(float64(width) * float64(origHeight) / float64(origWidth))
			} else if height > 0 {
				targetHeight = height
				targetWidth = int(float64(height) * float64(origWidth) / float64(origHeight))
			} else {
				// Neither width nor height is set, don't resize
				targetWidth = origWidth
				targetHeight = origHeight
			}

			// Prevent enlargement
			if targetWidth > origWidth || targetHeight > origHeight {
				targetWidth, targetHeight = origWidth, origHeight
			}

			logger.Info("Resizing image",
				slog.Int("width", width),
				slog.Int("height", height),
				slog.String("fileId", fileId),
				slog.String("contentType", contentType),
				slog.Int("quality", quality),
			)
			img = transform.Resize(img, targetWidth, targetHeight, transform.Lanczos)
		}

		buf := new(bytes.Buffer)

		switch contentType {
		case "image/jpeg", "image/jpg":
			err = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
		case "image/gif":
			err = gif.Encode(buf, img, nil)
		case "image/webp":
			err = webp.Encode(buf, img, &webp.Options{Quality: float32(quality)})
		default:
			// Fallback for formats without encoder (e.g. TIFF, BMP):
			// serve cached version as PNG
			w.Header().Set("Content-Type", "image/png")
			err = png.Encode(buf, img)
		}

		finalData := buf.Bytes()
		_ = os.WriteFile(cachePath, finalData, 0644)

		result := &imgResult{Data: finalData, ContentType: contentType}
		if len(finalData) < memCacheLimit {
			cs.imageDataCache.Add(hash, result.Data)
		}

		return result, nil
	})

	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	res := val.(*imgResult)
	w.Write(res.Data)
}

func (cs *ImageService) getCachePath(hash string) string {
	// Shard by the first 4 characters of the SHA1 hash
	subfolder := filepath.Join(hash[0:2], hash[2:4])
	fullDir := filepath.Join(cs.imageCacheFolder, subfolder)

	// Ensure the directories exist
	_ = os.MkdirAll(fullDir, 0755)

	return filepath.Join(fullDir, hash)
}
