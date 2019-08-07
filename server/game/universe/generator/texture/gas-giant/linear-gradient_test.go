package gasgiant_test

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"github.com/dayaftereh/discover/server/utils"

	gasgiant "github.com/dayaftereh/discover/server/game/universe/generator/texture/gas-giant"
	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
	"github.com/dayaftereh/discover/server/mathf"
)

func ColorToRGBA(c *textureset.Color) color.RGBA {
	r := uint8(255 * mathf.Clamp(c.R, 0.0, 1.0))
	g := uint8(255 * mathf.Clamp(c.G, 0.0, 1.0))
	b := uint8(255 * mathf.Clamp(c.B, 0.0, 1.0))
	a := uint8(255 * mathf.Clamp(c.A, 0.0, 1.0))

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

func WriteLinearGradientImage(t *testing.T, linearGradient *gasgiant.LinearGradient1D) {
	width, height := 1024, 200

	rect := image.Rect(0, 0, width, height)
	img := image.NewRGBA(rect)

	for x := 0; x < width; x++ {
		xn := float64(x) / float64(width)
		color := linearGradient.Color(xn)

		if color == nil {
			continue
		}

		rgb := ColorToRGBA(color)
		for y := 0; y < height; y++ {
			img.Set(x, y, rgb)
		}
	}

	fd, err := os.Create("linear-gradient.png")
	if err != nil {
		t.Fatal(err)
	}

	err = png.Encode(fd, img)
	if err != nil {
		t.Fatal(err)
	}
}

func RandomLinearGradient1D() *gasgiant.LinearGradient1D {
	linearGradient := gasgiant.NewLinearGradient1D()

	linearGradient.AddColor(0, textureset.NewRGBColor255(100.0, 100.0, 100.0))
	linearGradient.AddColor(0.15, textureset.NewRGBColor255(100.0, 100.0, 100.0))
	x := 0.25
	for x < 0.75 {
		r := utils.RandFloat64(0, 1)
		g := utils.RandFloat64(0, 1)
		b := utils.RandFloat64(0, 1)
		linearGradient.AddColor(x, textureset.NewRGBColor(r, g, b))
		x += utils.RandFloat64(0.001, 0.01)
	}
	linearGradient.AddColor(0.85, textureset.NewRGBColor255(150.0, 150.0, 150.0))
	linearGradient.AddColor(1.0, textureset.NewRGBColor255(150.0, 150.0, 150.0))

	return linearGradient
}

func TestLinearGradient1D(t *testing.T) {
	linearGradient := RandomLinearGradient1D()

	WriteLinearGradientImage(t, linearGradient)
}
