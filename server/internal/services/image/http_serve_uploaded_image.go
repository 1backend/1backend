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
	"github.com/gen2brain/avif"
	stdimage "image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
const transformCacheVersion = "v2-contain-default"

// @ID serveUploadedImage
// @Summary Serve Uploaded Image
// @Description Retrieves and serves a previously uploaded image file using its File ID.
// @Tags Image Svc
// @Accept json
// @Produce application/octet-stream
// @Param fileId path string true "FileID uniquely identifies the file itself (not an ID, which represents a specific replica)"
// @Param width query int false "Optional width to resize the image to"
// @Param height query int false "Optional height to resize the image to"
// @Param fit query string false "Resize strategy: contain|cover (default contain)"
// @Param position query string false "Crop anchor when fit=cover: center|top|bottom|left|right|top-left|top-right|bottom-left|bottom-right"
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

	params, err := parseImageParams(r.URL.Query())
	if err != nil {
		endpoint.WriteErr(w, http.StatusBadRequest, err)
		return
	}
	if r.URL.Query().Get("fit") == "" {
		params.Fit = fitContain
	}

	originalContentType, _ := cs.metaCache.Get(fileId)

	cacheKeyData := fmt.Sprintf("%s-%s-%d-%d-%d-%s-%s-%s",
		transformCacheVersion,
		fileId,
		params.Width,
		params.Height,
		params.Quality,
		params.RequestedFormat,
		params.Fit,
		params.Position,
	)

	targetContentType := "image/webp"
	if params.RequestedFormat != "" {
		// Normalize "jpg" to "jpeg" for consistent mime types
		fmtExt := strings.ToLower(params.RequestedFormat)
		if fmtExt == "jpg" {
			fmtExt = "jpeg"
		}
		targetContentType = "image/" + fmtExt
	}

	h := sha1.New()
	h.Write([]byte(cacheKeyData))
	hash := hex.EncodeToString(h.Sum(nil))

	var rsp *os.File

	if originalContentType == "" {
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

			originalContentType = hrsp.Header["Content-Type"][0]
			cs.metaCache.Add(fileId, originalContentType)

			return nil, nil
		})

		if err != nil {
			endpoint.WriteErr(w, http.StatusInternalServerError, err)
			return
		}

		defer rsp.Close()
	}

	if targetContentType != "" {
		w.Header().Set("Content-Type", targetContentType)
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
		logger.Info("Reading image from file service",
			slog.String("hash", hash),
			slog.String("fileId", fileId),
		)

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
		switch originalContentType {
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

		if params.Width > 0 || params.Height > 0 {
			logger.Info("Resizing image",
				slog.Int("width", params.Width),
				slog.Int("height", params.Height),
				slog.String("fileId", fileId),
				slog.String("originalContentType", originalContentType),
				slog.String("targetContentType", targetContentType),
				slog.Int("quality", params.Quality),
				slog.String("fit", params.Fit),
				slog.String("position", params.Position),
			)
			img = resizeWithFit(img, params.Width, params.Height, params.Fit, params.Position)
		}

		buf := new(bytes.Buffer)

		// Encode based on the target format
		switch targetContentType {
		case "image/jpeg", "image/jpg":
			err = jpeg.Encode(buf, img, &jpeg.Options{Quality: params.Quality})
		case "image/webp":
			// WebP is significantly smaller for photographic content (burgers!)
			err = webp.Encode(buf, img, &webp.Options{Quality: float32(params.Quality)})
		case "image/gif":
			err = gif.Encode(buf, img, nil)
		case "image/avif":
			err = avif.Encode(buf, img, avif.Options{Quality: params.Quality})
		case "image/png":
			// Use BestCompression for PNGs to shave off a few more KBs
			encoder := png.Encoder{CompressionLevel: png.BestCompression}
			err = encoder.Encode(buf, img)

		default:
			// If unknown, default to PNG as a safe web standard
			err = png.Encode(buf, img)
		}

		finalData := buf.Bytes()
		_ = os.WriteFile(cachePath, finalData, 0644)

		result := &imgResult{Data: finalData, ContentType: targetContentType}
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
