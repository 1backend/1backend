package imageservice

import (
	stdimage "image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResizeWithFitContainKeepsAspectRatio(t *testing.T) {
	t.Parallel()

	src := stdimage.NewRGBA(stdimage.Rect(0, 0, 200, 100))
	out := resizeWithFit(src, 120, 120, fitContain, positionCenter)

	require.Equal(t, 120, out.Bounds().Dx())
	require.Equal(t, 60, out.Bounds().Dy())
}

func TestResizeWithFitCoverAppliesPosition(t *testing.T) {
	t.Parallel()

	src := stdimage.NewRGBA(stdimage.Rect(0, 0, 200, 100))
	fillRect(src, stdimage.Rect(0, 0, 100, 100), color.RGBA{255, 0, 0, 255})   // left red
	fillRect(src, stdimage.Rect(100, 0, 200, 100), color.RGBA{0, 255, 0, 255}) // right green

	out := resizeWithFit(src, 50, 50, fitCover, positionLeft)
	require.Equal(t, color.RGBA{255, 0, 0, 255}, rgbaAt(out, 1, 1))

	out = resizeWithFit(src, 50, 50, fitCover, positionRight)
	require.Equal(t, color.RGBA{0, 255, 0, 255}, rgbaAt(out, 1, 1))
}

func TestResizeWithFitCoverTallImageProducesRequestedSize(t *testing.T) {
	t.Parallel()

	src := stdimage.NewRGBA(stdimage.Rect(0, 0, 2000, 20000))
	out := resizeWithFit(src, 1200, 630, fitCover, positionCenter)

	require.Equal(t, 1200, out.Bounds().Dx())
	require.Equal(t, 630, out.Bounds().Dy())
}

func TestParseImageParams(t *testing.T) {
	t.Parallel()

	params, err := parseImageParams(map[string][]string{
		"width":    {"1200"},
		"height":   {"630"},
		"fit":      {"cover"},
		"position": {"top-right"},
	})
	require.NoError(t, err)
	require.Equal(t, 1200, params.Width)
	require.Equal(t, 630, params.Height)
	require.Equal(t, fitCover, params.Fit)
	require.Equal(t, positionTopRight, params.Position)

	_, err = parseImageParams(map[string][]string{"fit": {"stretch"}})
	require.Error(t, err)
}

func TestParseImageParamsDefaultsFitToContainWhenFitNotProvided(t *testing.T) {
	t.Parallel()

	params, err := parseImageParams(map[string][]string{
		"width":  {"360"},
		"height": {"100"},
	})
	require.NoError(t, err)
	require.Equal(t, fitContain, params.Fit)
}

func fillRect(img *stdimage.RGBA, rect stdimage.Rectangle, c color.RGBA) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.SetRGBA(x, y, c)
		}
	}
}

func rgbaAt(img stdimage.Image, x, y int) color.RGBA {
	r, g, b, a := img.At(x, y).RGBA()
	return color.RGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}
