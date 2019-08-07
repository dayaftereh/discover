package gasgiant

import (
	"image"
	"image/color"
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
	// noise between [-1.0, 1.0]
	noiseSimplex := noise.NewImplicitSimplex(gasGiantGenerator.seed)
	noiseRanged := noise.NewImplicitRanged(noiseSimplex, -1.0, 1.0)

	// fractal noise
	gasNoiseFractal := noise.NewImplicitFractalNormalized(noiseRanged, gasNoiseOctaves, gasNoiseFrequency, gasNoisePersistence)
	gasNoiseRidgedFractal := noise.NewImplicitFractalRidged(noiseRanged, ridgedOctaves, ridgedFrequency, ridgedPersistence)

	// storm noise
	stormNoise := noise.NewImplicitInputTransform(noiseRanged, func(value float64) float64 {
		return value * 0.1
	})
	stormNoise1 := noise.NewImplicitInputTransform(noiseRanged, func(value float64) float64 {
		return value * 2.0
	})
	stormNoise2 := noise.NewImplicitInputTransform(noiseRanged, func(value float64) float64 {
		return (value + 800.0) * 2.0
	})
	stormNoise3 := noise.NewImplicitInputTransform(noiseRanged, func(value float64) float64 {
		return (value + 1600.0) * 2.0
	})

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

			// Base noise
			n1 := gasNoiseFractal.Get3D(x1, y1, z1) * 0.01
			n2 := gasNoiseRidgedFractal.Get3D(x1, y1, z1)*0.015 - 0.01

			// Get the three threshold samples
			s := 0.6
			t1 := stormNoise1.Get3D(x1, y1, z1) - s
			t2 := stormNoise2.Get3D(x1, y1, z1) - s
			t3 := stormNoise3.Get3D(x1, y1, z1) - s
			// Intersect them and get rid of negatives
			threshold := math.Max(t1*t2*t3, 0.0)

			// Storms
			n3 := stormNoise.Get3D(x1, y1, z1) * threshold * 10.0

			// cumulate
			n := n1 + n2 + n3

			gasGiantGenerator.fractalData.Set(x, y, n)

			curLon += xDelta
		}

		curLat += yDelta
	}

}

func (gasGiantGenerator *GasGiantGenerator) latLonToXYZ(lat, lon float64) (float64, float64, float64) {
	r := math.Cos(mathf.ToRadians(lon))
	x := r * math.Cos(mathf.ToRadians(lat))
	y := math.Sin(mathf.ToRadians(lon))
	z := r * math.Sin(mathf.ToRadians(lat))

	return x, y, z
}

func (gasGiantGenerator *GasGiantGenerator) GetFractalImage() *image.RGBA {
	rect := image.Rect(0, 0, gasGiantGenerator.width*5, gasGiantGenerator.height)
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
