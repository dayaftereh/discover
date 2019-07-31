package noise_test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
	"github.com/dayaftereh/discover/server/utils"
)

func WriteImage(t *testing.T, implicitBase noise.ImplicitBase, name string) {
	width, height := 1024, 1024
	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx := float64(x) / float64(width)
			ny := float64(y) / float64(height)
			value := implicitBase.Get2D(nx, ny)
			cx := uint8(255.0 * value)
			gray := color.NRGBA{cx, cx, cx, 0xff}
			img.Set(x, y, gray)
		}
	}

	filename := fmt.Sprintf("./%s.png", name)
	f, err := os.Create(filename)

	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplicitFractalGetD2(t *testing.T) {
	seed := utils.RandInt64(0, 100)
	t.Logf("seed: %d", seed)
	frequency, octaves := 0.15, int64(4)
	implicitFractal := noise.NewImplicitFractal(frequency, octaves, true, seed)

	WriteImage(t, implicitFractal, "implicit-fractal-d2")
}
