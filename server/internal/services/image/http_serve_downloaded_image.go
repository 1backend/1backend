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
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/chai2010/webp"
	"github.com/gen2brain/avif"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"

	"github.com/1backend/1backend/sdk/go/client"
	"github.com/1backend/1backend/sdk/go/endpoint"
	"github.com/1backend/1backend/sdk/go/logger"

	openapi "github.com/1backend/1backend/clients/go"
)

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

	width, height, quality, requestedFormat, err := parseImageParams(r)
	if err != nil {
		endpoint.WriteErr(w, http.StatusBadRequest, err)
		return
	}

	originalContentType, _ := cs.metaCache.Get("download:" + rawURL)

	targetContentType := buildTargetContentType(requestedFormat)
	cacheKeyData := fmt.Sprintf("download:%s-%d-%d-%d-%s", rawURL, width, height, quality, requestedFormat)
	hash := sha1Hex(cacheKeyData)

	if targetContentType != "" {
		w.Header().Set("Content-Type", targetContentType)
	}
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")

	// RAM cache
	if data, ok := cs.imageDataCache.Get(hash); ok {
		_, _ = w.Write(data)
		return
	}

	// Disk cache
	cachePath := cs.getCachePath(hash)
	if data, err := os.ReadFile(cachePath); err == nil {
		if len(data) < memCacheLimit {
			cs.imageDataCache.Add(hash, data)
		}
		_, _ = w.Write(data)
		return
	}

	val, err, _ := cs.sf.Do(hash, func() (interface{}, error) {
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

		result, err := cs.processImage(
			rsp,
			originalContentType,
			targetContentType,
			width,
			height,
			quality,
			cachePath,
		)
		if err != nil {
			return nil, err
		}

		return result, nil
	})

	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			status = http.StatusNotFound
		}
		endpoint.WriteErr(w, status, err)
		return
	}

	res := val.(*imgResult)
	_, _ = w.Write(res.Data)
}

func parseImageParams(r *http.Request) (int, int, int, string, error) {
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	qualityStr := r.URL.Query().Get("quality")
	requestedFormat := r.URL.Query().Get("format")

	width := 0
	height := 0
	quality := 85

	var err error
	if widthStr != "" {
		widthStr = strings.TrimSuffix(widthStr, "px")
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			return 0, 0, 0, "", errors.New("invalid width")
		}
	}
	if heightStr != "" {
		heightStr = strings.TrimSuffix(heightStr, "px")
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			return 0, 0, 0, "", errors.New("invalid height")
		}
	}
	if qualityStr != "" {
		quality, err = strconv.Atoi(qualityStr)
		if err != nil {
			return 0, 0, 0, "", errors.New("invalid quality")
		}
	}

	return width, height, quality, requestedFormat, nil
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
		img = resizePreservingAspect(img, width, height)
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

func resizePreservingAspect(img stdimage.Image, width, height int) stdimage.Image {
	bounds := img.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	targetWidth := origWidth
	targetHeight := origHeight

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
	}

	if targetWidth > origWidth || targetHeight > origHeight {
		targetWidth = origWidth
		targetHeight = origHeight
	}

	return transform.Resize(img, targetWidth, targetHeight, transform.Lanczos)
}

func (cs *ImageService) openDownloadedFile(
	ctx context.Context,
	rawURL string,
) (*os.File, *http.Response, error) {
	api := cs.options.ClientFactory.Client(client.WithToken(cs.token)).FileSvcAPI

	// First try serve existing download
	file, rsp, err := api.ServeDownload(ctx, rawURL).Execute()
	if err == nil {
		return file, rsp, nil
	}

	// If not found, trigger download
	var apiErr openapi.GenericOpenAPIError
	if errors.As(err, &apiErr) && rsp != nil && rsp.StatusCode == http.StatusNotFound {
		_, _, derr := api.DownloadFile(ctx).
			Body(openapi.FileSvcDownloadFileRequest{
				Url: rawURL,
			}).
			Execute()
		if derr != nil {
			return nil, nil, errors.Wrap(derr, "failed to download source file")
		}

		file, rsp, err = api.ServeDownload(ctx, rawURL).Execute()
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to serve downloaded file after download")
		}
		return file, rsp, nil
	}

	return nil, nil, errors.Wrap(err, "failed to serve downloaded file")
}
