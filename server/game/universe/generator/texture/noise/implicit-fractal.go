package noise

import (
	"math"

	"github.com/ojrac/opensimplex-go"
)

// Implementaion based on https://github.com/jongallant/WorldGeneratorFinal/blob/master/Assets/Scripts/AccidentalNoise/Implicit/ImplicitFractal.cs

type ImplicitFractal struct {
	noise opensimplex.Noise

	octaves    int64
	h          float64
	lacunarity float64
	frequency  float64
}

// NewImplicitFractal creates a new ImplicitFractal to abstract the noise generation
// frequency specifies the density of the function.
// octaves determines how many layers contribute to the fractal
// normalized returns values between [0,1]
func NewImplicitFractal(frequency float64, octaves int64, normalized bool, seed int64) *ImplicitFractal {
	var noise opensimplex.Noise
	if normalized {
		noise = opensimplex.NewNormalized(seed)
	} else {
		noise = opensimplex.New(seed)
	}

	return &ImplicitFractal{
		noise:      noise,
		lacunarity: 2.0,
		h:          1.0,
		octaves:    octaves,
		frequency:  frequency,
	}
}

func (implicitFractal *ImplicitFractal) Get2D(x, y float64) float64 {
	value := 1.0

	x *= implicitFractal.frequency
	y *= implicitFractal.frequency

	minValue, maxValue := 1.0, 1.0
	for i := int64(0); i < implicitFractal.octaves; i++ {
		exp := math.Pow(implicitFractal.lacunarity, float64(-i)*implicitFractal.h)

		minValue *= -1.0*exp + 1.0
		maxValue *= 1.0*exp + 1.0

		value *= implicitFractal.noise.Eval2(x, y)*exp + 1.0

		x *= implicitFractal.lacunarity
		y *= implicitFractal.lacunarity
	}

	scale := 2.0 / (maxValue - minValue)
	bias := -1.0 - minValue*scale

	return value*scale + bias
}

func (implicitFractal *ImplicitFractal) Get3D(x, y, z float64) float64 {
	value := 1.0

	x *= implicitFractal.frequency
	y *= implicitFractal.frequency
	z *= implicitFractal.frequency

	minValue, maxValue := 1.0, 1.0
	for i := int64(0); i < implicitFractal.octaves; i++ {
		exp := math.Pow(implicitFractal.lacunarity, float64(-i)*implicitFractal.h)

		minValue *= -1.0*exp + 1.0
		maxValue *= 1.0*exp + 1.0

		value *= implicitFractal.noise.Eval3(x, y, z)*exp + 1.0

		x *= implicitFractal.lacunarity
		y *= implicitFractal.lacunarity
		z *= implicitFractal.lacunarity
	}

	scale := 2.0 / (maxValue - minValue)
	bias := -1.0 - minValue*scale

	return value*scale + bias
}

func (implicitFractal *ImplicitFractal) Get4D(x, y, z, w float64) float64 {
	value := 1.0

	x *= implicitFractal.frequency
	y *= implicitFractal.frequency
	z *= implicitFractal.frequency
	w *= implicitFractal.frequency

	minValue, maxValue := 1.0, 1.0
	for i := int64(0); i < implicitFractal.octaves; i++ {
		exp := math.Pow(implicitFractal.lacunarity, float64(-i)*implicitFractal.h)

		minValue *= -1.0*exp + 1.0
		maxValue *= 1.0*exp + 1.0

		value *= implicitFractal.noise.Eval4(x, y, z, w)*exp + 1.0

		x *= implicitFractal.lacunarity
		y *= implicitFractal.lacunarity
		z *= implicitFractal.lacunarity
		w *= implicitFractal.lacunarity
	}

	scale := 2.0 / (maxValue - minValue)
	bias := -1.0 - minValue*scale

	return value*scale + bias
}
