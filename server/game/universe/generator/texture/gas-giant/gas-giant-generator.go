package gasgiant

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture"
	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
	"github.com/dayaftereh/discover/server/mathf"
)

type GasGiantGenerator struct {
	seed           int64
	width          int
	height         int
	linearGradient *LinearGradient1D
	fractalData    *texture.MapData
}

func NewGasGiantGenerator(linearGradient *LinearGradient1D, width, height int, seed int64) *GasGiantGenerator {
	return &GasGiantGenerator{
		seed:           seed,
		width:          width,
		height:         height,
		linearGradient: linearGradient,
	}
}

func (gasGiantGenerator *GasGiantGenerator) Init() {
	fractalFrequency := 10.0
	fractalOctaves := int64(6)
	fractalPersistence := 0.8
	fractalMap := noise.NewImplicitFractal2(fractalOctaves, fractalFrequency, fractalPersistence, gasGiantGenerator.seed)

	ridgedMap := noise.NewImplicitFractalRidged(5, 5.8, 0.75, gasGiantGenerator.seed)

	gasGiantGenerator.fractalData = texture.NewMapData(gasGiantGenerator.width, gasGiantGenerator.height)

	southLatBound := -180.0
	northLatBound := 180.0
	westLonBound := -90.0
	eastLonBound := 90.0

	lonExtent := eastLonBound - westLonBound
	latExtent := northLatBound - southLatBound

	xDelta := lonExtent / float64(gasGiantGenerator.width)
	yDelta := latExtent / float64(gasGiantGenerator.height)

	curLon := westLonBound
	curLat := southLatBound

	for x := 0; x < gasGiantGenerator.width; x++ {
		curLon = westLonBound
		for y := 0; y < gasGiantGenerator.height; y++ {
			x1, y1, z1 := gasGiantGenerator.latLonToXYZ(curLat, curLon)

			fractalValue := fractalMap.Get3D(x1, y1, z1) * 0.01
			ridgedValue := ridgedMap.Get3D(x1, y1, z1)*0.015 - 0.01
			gasGiantGenerator.fractalData.Set(x, y, fractalValue+ridgedValue)

			curLon += xDelta
		}

		curLat += yDelta
	}

	log.Println(gasGiantGenerator.fractalData.MinValue, gasGiantGenerator.fractalData.MaxValue)
}

func (gasGiantGenerator *GasGiantGenerator) latLonToXYZ(lat, lon float64) (float64, float64, float64) {
	r := math.Cos(mathf.ToRadians(lon))
	x := r * math.Cos(mathf.ToRadians(lat))
	y := math.Sin(mathf.ToRadians(lon))
	z := r * math.Sin(mathf.ToRadians(lat))

	return x, y, z
}

func (gasGiantGenerator *GasGiantGenerator) GetFractalImage() *image.RGBA {
	rect := image.Rect(0, 0, gasGiantGenerator.width, gasGiantGenerator.height)
	img := image.NewRGBA(rect)

	for x := 0; x < gasGiantGenerator.width; x++ {
		for y := 0; y < gasGiantGenerator.height; y++ {
			yn := float64(y) / float64(gasGiantGenerator.height)
			value := yn + gasGiantGenerator.fractalData.Get(x, y)
			color := gasGiantGenerator.linearGradient.Color(value)

			rgb := gasGiantGenerator.color2RGBA(color)
			img.Set(x, y, rgb)
		}
	}

	return img
}

func (gasGiantGenerator *GasGiantGenerator) color2RGBA(c *textureset.Color) color.RGBA {
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
