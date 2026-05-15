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
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"github.com/gen2brain/avif"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	image "github.com/1backend/1backend/server/internal/services/image/types"
)

type ErrResponse image.ErrorResponse

// @ID serveDownloadedImage
// @Summary Serve Downloaded Image
// @Description Retrieves, caches, resizes, and serves an image referenced by its original URL.
// @Tags Image Svc
// @Accept json
// @Produce application/octet-stream
// @Param url path string true "Original URL of the downloaded file (path-escaped)"
// @Param width query int false "Optional width to resize the image to"
// @Param height query int false "Optional height to resize the image to"
// @Param quality query int false "Optional quality for lossy output formats (default 85)"
// @Param format query string false "Optional output format: webp, jpeg, png, gif, avif"
// @Param fit query string false "Resize strategy: contain|cover (default contain)"
// @Param position query string false "Crop anchor when fit=cover: center|top|bottom|left|right|top-left|top-right|bottom-left|bottom-right"
// @Success 200 {file} binary "Image served successfully"
// @Failure 400 {object} image.ErrorResponse "Invalid URL"
// @Failure 404 {object} image.ErrorResponse "File Not Found"
// @Failure 500 {object} image.ErrorResponse "Internal Server Error"
// @Router /image-svc/serve/download/{url} [get]
func (cs *ImageService) ServeDownloadedImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	rawURL, err := url.PathUnescape(vars["url"])
	if err != nil || rawURL == "" {
		endpoint.WriteString(w, http.StatusBadRequest, "Invalid download URL")
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
	ctx, span := startImageSpan(r.Context(), "image.serve_downloaded",
		attribute.Int("image.width", params.Width),
		attribute.Int("image.height", params.Height),
		attribute.String("image.fit", params.Fit),
		attribute.String("image.format", params.RequestedFormat),
	)
	defer span.End()
	r = r.WithContext(ctx)

	originalContentType, _ := cs.metaCache.Get("download:" + rawURL)

	targetContentType := buildTargetContentType(params.RequestedFormat)
	cacheKeyData := fmt.Sprintf("%s-download:%s-%d-%d-%d-%s-%s-%s",
		transformCacheVersion,
		rawURL,
		params.Width,
		params.Height,
		params.Quality,
		params.RequestedFormat,
		params.Fit,
		params.Position,
	)
	hash := sha1Hex(cacheKeyData)

	// RAM cache
	if data, ok := cs.imageDataCache.Get(hash); ok {
		recordImageServe(r.Context(), "download", "ram", "success", int64(len(data)), params)
		writeCachedImage(w, targetContentType, data)
		return
	}

	// Disk cache
	cachePath := cs.getCachePath(hash)
	if data, err := os.ReadFile(cachePath); err == nil {
		if len(data) < memCacheLimit {
			cs.imageDataCache.Add(hash, data)
		}
		recordImageServe(r.Context(), "download", "disk", "success", int64(len(data)), params)
		writeCachedImage(w, targetContentType, data)
		return
	}

	didTransform := false
	transformCache := "cold_miss"
	val, err, _ := cs.sf.Do(hash, func() (interface{}, error) {
		transformStart := time.Now()
		logger.Info("Reading image from file service download",
			slog.String("hash", hash),
			slog.String("url", rawURL),
		)

		rsp, hrsp, err := cs.openDownloadedFile(r.Context(), rawURL)
		if err != nil {
			return nil, err
		}
		defer rsp.Close()

		if originalContentType == "" {
			originalContentType = hrsp.Header.Get("Content-Type")
			cs.metaCache.Add("download:"+rawURL, originalContentType)
		}
		transformCache = imageCacheLabelFromStorageSource(hrsp.Header.Get(imageStorageSourceHeader))

		result, err := cs.processImage(
			rsp,
			originalContentType,
			targetContentType,
			params.Width,
			params.Height,
			params.Quality,
			params.Fit,
			params.Position,
			cachePath,
		)
		if err != nil {
			recordImageTransform(r.Context(), "download", "error", time.Since(transformStart))
			return nil, err
		}
		recordImageTransform(r.Context(), "download", "success", time.Since(transformStart))
		didTransform = true

		return result, nil
	})

	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			status = http.StatusNotFound
		}
		logger.Error("Error download file", slog.Any("error", err))
		span.RecordError(err)
		span.SetStatus(codes.Error, "transform")
		recordImageServe(r.Context(), "download", "cold_miss", "error", -1, params)
		endpoint.WriteErr(w, status, err)
		return
	}

	res := val.(*imgResult)
	cache := "coalesced"
	if didTransform {
		cache = transformCache
	}
	recordImageServe(r.Context(), "download", cache, "success", int64(len(res.Data)), params)
	writeCachedImage(w, res.ContentType, res.Data)
}

func writeCachedImage(w http.ResponseWriter, contentType string, data []byte) {
	if contentType != "" {
		w.Header().Set("Content-Type", contentType)
	}
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	_, _ = w.Write(data)
}

func buildTargetContentType(requestedFormat string) string {
	targetContentType := "image/webp"
	if requestedFormat == "" {
		return targetContentType
	}

	fmtExt := strings.ToLower(requestedFormat)
	if fmtExt == "jpg" {
		fmtExt = "jpeg"
	}
	return "image/" + fmtExt
}

func sha1Hex(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (cs *ImageService) processImage(
	src io.Reader,
	originalContentType string,
	targetContentType string,
	width int,
	height int,
	quality int,
	fit string,
	position string,
	cachePath string,
) (*imgResult, error) {
	var (
		img stdimage.Image
		err error
	)

	switch originalContentType {
	case "image/png":
		img, err = png.Decode(src)
	case "image/jpeg", "image/jpg":
		img, err = jpeg.Decode(src)
	case "image/gif":
		img, err = gif.Decode(src)
	case "image/webp":
		img, err = webp.Decode(src)
	case "image/tiff":
		img, err = tiff.Decode(src)
	case "image/bmp":
		img, err = bmp.Decode(src)
	case "image/avif":
		img, err = avif.Decode(src)
	default:
		img, _, err = stdimage.Decode(src)
	}
	if err != nil {
		return nil, errors.Wrap(err, "decode err")
	}

	if width > 0 || height > 0 {
		img = resizeWithFit(img, width, height, fit, position)
	}

	buf := new(bytes.Buffer)

	switch targetContentType {
	case "image/jpeg", "image/jpg":
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
	case "image/webp":
		err = webp.Encode(buf, img, &webp.Options{Quality: float32(quality)})
	case "image/gif":
		err = gif.Encode(buf, img, nil)
	case "image/avif":
		err = avif.Encode(buf, img, avif.Options{Quality: quality})
	case "image/png":
		encoder := png.Encoder{CompressionLevel: png.BestCompression}
		err = encoder.Encode(buf, img)
	default:
		err = png.Encode(buf, img)
	}

	if err != nil {
		return nil, errors.Wrap(err, "encode err")
	}

	finalData := buf.Bytes()
	_ = os.WriteFile(cachePath, finalData, 0644)

	result := &imgResult{
		Data:        finalData,
		ContentType: targetContentType,
	}
	if len(finalData) < memCacheLimit {
		cs.imageDataCache.Add(filepath.Base(cachePath), result.Data)
	}

	return result, nil
}

func (cs *ImageService) openDownloadedFile(
	ctx context.Context,
	rawURL string,
) (*os.File, *http.Response, error) {
	api := cs.options.ClientFactory.Client(client.WithToken(cs.token)).FileSvcAPI

	file, rsp, err := api.ServeDownload(ctx, rawURL).Execute()
	if err == nil {
		recordImageFileOpen(ctx, "download", "success")
		return file, rsp, nil
	}

	if rsp != nil && rsp.StatusCode == http.StatusNotFound {
		recordImageFileOpen(ctx, "download", "not_found_retry")
		file, rsp, err = api.ServeDownload(ctx, rawURL).Execute()
		if err != nil {
			logger.Error(
				"Failed to serve file",
				slog.Any("error", err),
				slog.Any("url", rawURL),
			)
			recordImageFileOpen(ctx, "download", "retry_error")
			return nil, nil, errors.Wrap(err, "failed to serve downloaded file")
		}
		recordImageFileOpen(ctx, "download", "retry_success")
		return file, rsp, nil
	}

	recordImageFileOpen(ctx, "download", "error")
	return nil, rsp, errors.Wrap(err, "failed to serve downloaded file")
}
