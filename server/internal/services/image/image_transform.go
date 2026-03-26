/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package imageservice

import (
	stdimage "image"
	"image/draw"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/transform"
	"github.com/pkg/errors"
)

const (
	fitContain = "contain"
	fitCover   = "cover"
)

const (
	positionCenter      = "center"
	positionTop         = "top"
	positionBottom      = "bottom"
	positionLeft        = "left"
	positionRight       = "right"
	positionTopLeft     = "top-left"
	positionTopRight    = "top-right"
	positionBottomLeft  = "bottom-left"
	positionBottomRight = "bottom-right"
)

type imageParams struct {
	Width           int
	Height          int
	Quality         int
	RequestedFormat string
	Fit             string
	Position        string
}

func parseImageParams(query map[string][]string) (*imageParams, error) {
	width, err := parsePxOrInt(firstQueryValue(query, "width"))
	if err != nil {
		return nil, errors.New("invalid width")
	}
	height, err := parsePxOrInt(firstQueryValue(query, "height"))
	if err != nil {
		return nil, errors.New("invalid height")
	}

	quality := 85
	qualityStr := firstQueryValue(query, "quality")
	if qualityStr != "" {
		quality, err = strconv.Atoi(qualityStr)
		if err != nil {
			return nil, errors.New("invalid quality")
		}
	}

	fit := strings.ToLower(firstQueryValue(query, "fit"))
	if fit == "" {
		fit = fitContain
	}
	if fit != fitContain && fit != fitCover {
		return nil, errors.New("invalid fit")
	}

	position := strings.ToLower(firstQueryValue(query, "position"))
	if position == "" {
		position = positionCenter
	}
	switch position {
	case positionCenter, positionTop, positionBottom, positionLeft, positionRight,
		positionTopLeft, positionTopRight, positionBottomLeft, positionBottomRight:
	default:
		return nil, errors.New("invalid position")
	}

	return &imageParams{
		Width:           width,
		Height:          height,
		Quality:         quality,
		RequestedFormat: firstQueryValue(query, "format"),
		Fit:             fit,
		Position:        position,
	}, nil
}

func firstQueryValue(query map[string][]string, key string) string {
	if values, ok := query[key]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}

func parsePxOrInt(v string) (int, error) {
	if v == "" {
		return 0, nil
	}
	v = strings.TrimSuffix(v, "px")
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, errors.New("must be non-negative")
	}
	return n, nil
}

func resizeWithFit(img stdimage.Image, width, height int, fit, position string) stdimage.Image {
	if width == 0 && height == 0 {
		return img
	}

	if fit == fitCover && width > 0 && height > 0 {
		return resizeCover(img, width, height, position)
	}

	return resizeContain(img, width, height)
}

func resizeContain(img stdimage.Image, width, height int) stdimage.Image {
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

func resizeCover(img stdimage.Image, width, height int, position string) stdimage.Image {
	b := img.Bounds()
	origWidth := b.Dx()
	origHeight := b.Dy()

	scaleX := float64(width) / float64(origWidth)
	scaleY := float64(height) / float64(origHeight)
	scale := scaleX
	if scaleY > scale {
		scale = scaleY
	}

	if scale > 1 {
		scale = 1
	}

	resizedWidth := int(float64(origWidth) * scale)
	resizedHeight := int(float64(origHeight) * scale)
	if resizedWidth < 1 {
		resizedWidth = 1
	}
	if resizedHeight < 1 {
		resizedHeight = 1
	}

	resized := transform.Resize(img, resizedWidth, resizedHeight, transform.Lanczos)

	cropWidth := width
	cropHeight := height
	if cropWidth > resizedWidth {
		cropWidth = resizedWidth
	}
	if cropHeight > resizedHeight {
		cropHeight = resizedHeight
	}

	x, y := coverCropOrigin(resizedWidth, resizedHeight, cropWidth, cropHeight, position)
	return cropImage(resized, x, y, cropWidth, cropHeight)
}

func coverCropOrigin(imgW, imgH, cropW, cropH int, position string) (int, int) {
	centerX := (imgW - cropW) / 2
	centerY := (imgH - cropH) / 2

	switch position {
	case positionTop:
		return centerX, 0
	case positionBottom:
		return centerX, imgH - cropH
	case positionLeft:
		return 0, centerY
	case positionRight:
		return imgW - cropW, centerY
	case positionTopLeft:
		return 0, 0
	case positionTopRight:
		return imgW - cropW, 0
	case positionBottomLeft:
		return 0, imgH - cropH
	case positionBottomRight:
		return imgW - cropW, imgH - cropH
	default:
		return centerX, centerY
	}
}

func cropImage(img stdimage.Image, x, y, w, h int) stdimage.Image {
	dst := stdimage.NewRGBA(stdimage.Rect(0, 0, w, h))
	srcRect := stdimage.Rect(x, y, x+w, y+h)
	draw.Draw(dst, dst.Bounds(), img, srcRect.Min, draw.Src)
	return dst
}
