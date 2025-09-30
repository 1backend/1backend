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
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

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

	rsp, hrsp, err := cs.options.ClientFactory.Client(client.WithToken(cs.token)).
		FileSvcAPI.ServeUpload(context.Background(), fileId).
		Execute()
	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
	}

	contentType := "image/png"
	if len(hrsp.Header["Content-Type"]) > 0 {
		contentType = hrsp.Header.Get("Content-Type")
	}

	resampleFilter := transform.Lanczos

	cacheKeyData := fmt.Sprintf("%s-%d-%d-%d-%s-%v",
		fileId,
		width,
		height,
		quality,
		contentType,
		resampleFilter,
	)

	h := sha1.New()
	h.Write([]byte(cacheKeyData))
	hash := hex.EncodeToString(h.Sum(nil))
	cachePath := filepath.Join(cs.imageCacheFolder, hash)

	if f, err := os.Open(cachePath); err == nil {
		defer f.Close()
		switch contentType {
		case "image/jpeg", "image/jpg":
			w.Header().Set("Content-Type", "image/jpeg")
		case "image/gif":
			w.Header().Set("Content-Type", "image/gif")
		default:
			w.Header().Set("Content-Type", "image/png")
		}
		io.Copy(w, f)
		return
	}

	img, _, err := stdimage.Decode(rsp)
	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
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

		logger.Info("Resizing image",
			slog.Int("width", width),
			slog.Int("height", height),
			slog.String("fileId", fileId),
			slog.String("contentType", contentType),
			slog.Int("quality", quality),
		)
		img = transform.Resize(img, targetWidth, targetHeight, resampleFilter)
	}

	outFile, err := os.Create(cachePath)
	if err != nil {
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
		return
	}
	defer outFile.Close()

	switch contentType {
	case "image/jpeg", "image/jpg":
		w.Header().Set("Content-Type", "image/jpeg")
		err = jpeg.Encode(io.MultiWriter(w, outFile), img, &jpeg.Options{Quality: quality})
	case "image/gif":
		w.Header().Set("Content-Type", "image/gif")
		err = gif.Encode(io.MultiWriter(w, outFile), img, nil)
	default:
		w.Header().Set("Content-Type", "image/png")
		err = png.Encode(io.MultiWriter(w, outFile), img)
	}

	switch {
	case err == nil:
	case errors.Is(err, stdimage.ErrFormat):
		endpoint.WriteErr(w, http.StatusUnsupportedMediaType, errors.New("unsupported image format"))
	default:
		endpoint.WriteErr(w, http.StatusInternalServerError, err)
	}
}
