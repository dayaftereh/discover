package noise_test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
	"github.com/dayaftereh/discover/server/utils"
)

func FuncWriteImage(t *testing.T, width, height int, name string, fn func(x, y float64) float64) {
	minValue := math.Inf(1)
	maxValue := math.Inf(-1)
	values := make([]float64, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			nx := float64(x) / float64(width)
			ny := float64(y) / float64(height)
			value := fn(nx, ny)

			index := y*width + x
			values[index] = value

			minValue = math.Min(value, minValue)
			maxValue = math.Max(value, maxValue)
		}
	}

	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			value := values[index]

			value = (value - minValue) / (maxValue - minValue)

			cx := uint8(255.0 * value)
			gray := color.RGBA{cx, cx, cx, 0xff}
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

func WriteImage2D(t *testing.T, implicitBase noise.ImplicitBase, name string) {
	FuncWriteImage(t, 1024, 1024, name, func(x, y float64) float64 {
		return implicitBase.Get2D(x, y)
	})
}

func WriteImage4D(t *testing.T, implicitBase noise.ImplicitBase, name string) {
	FuncWriteImage(t, 1024, 1024, name, func(x, y float64) float64 {
		dx := 2.0
		dy := 2.0

		nx := math.Cos(x*2.0*math.Pi) * dx / (2.0 * math.Pi)
		ny := math.Cos(y*2.0*math.Pi) * dy / (2.0 * math.Pi)
		nz := math.Sin(x*2.0*math.Pi) * dx / (2.0 * math.Pi)
		nw := math.Sin(y*2.0*math.Pi) * dy / (2.0 * math.Pi)

		value := implicitBase.Get4D(nx, ny, nz, nw)

		return value
	})
}

func TestImplicitFractalGetD2(t *testing.T) {
	seed := utils.RandInt64(0, 100)
	t.Logf("seed: %d", seed)
	frequency, octaves := 1.15, int64(4)
	implicitFractal := noise.NewImplicitFractal(frequency, octaves, true, seed)

	WriteImage2D(t, implicitFractal, "implicit-fractal-2d")
}

func TestImplicitFractalGetD4(t *testing.T) {
	seed := utils.RandInt64(0, 100)
	t.Logf("seed: %d", seed)
	frequency, octaves := 1.15, int64(4)
	implicitFractal := noise.NewImplicitFractal(frequency, octaves, true, seed)

	WriteImage4D(t, implicitFractal, "implicit-fractal-4d")
}
