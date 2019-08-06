package noise

import "github.com/ojrac/opensimplex-go"

type ImplicitFractal2 struct {
	noise opensimplex.Noise

	/*
		Octaves is the number of iterations, or the depth of the fractal. The more octaves, the better quality, but performance suffers. Keep this number as small as you can.
		Frequency determines the wavelength of your noise. A low frequency means a few wide hills, and a high frequency means many skinny hills. Each successive octave doubles the frequency, which is what gives us the 4 hills on top of 1 in the above example.
		Persistence determines how much each successive octave affects the end result. If it's 1.0, then every octave holds the same weight. If it's 0.0, then only the first octave does anything. At 0.5, each successive octave applies half as much weight to the end product. A good value is typically in the range of 0.8, depending on what you are trying to make.
	*/

	octaves     int64
	frequency   float64
	persistence float64
}

func NewImplicitFractal2(octaves int64, frequency, persistence float64, seed int64) *ImplicitFractal2 {
	return &ImplicitFractal2{
		noise:       opensimplex.NewNormalized(seed),
		octaves:     octaves,
		frequency:   frequency,
		persistence: persistence,
	}
}

func (implicitFractal *ImplicitFractal2) execute(fn func(frequency float64) float64) float64 {
	// Total value so far
	total := 0.0
	// Accumulates highest theoretical amplitude
	maxAmplitude := 0.0

	amplitude := 1.0
	frequency := implicitFractal.frequency
	for i := int64(0); i < implicitFractal.octaves; i++ {
		// Get the noise sample
		total += fn(frequency) * amplitude

		// Make the wavelength twice as small
		frequency *= 2.0
		// Add to our maximum possible amplitude
		maxAmplitude += amplitude

		// Reduce amplitude according to persistence for the next octave
		amplitude *= implicitFractal.persistence
	}

	// Scale the result by the maximum amplitude
	return total / maxAmplitude
}

func (implicitFractal *ImplicitFractal2) Get2D(x, y float64) float64 {
	return implicitFractal.execute(func(frequency float64) float64 {
		return implicitFractal.noise.Eval2(x*frequency, y*frequency)
	})
}

func (implicitFractal *ImplicitFractal2) Get3D(x, y, z float64) float64 {
	return implicitFractal.execute(func(frequency float64) float64 {
		return implicitFractal.noise.Eval3(x*frequency, y*frequency, z*frequency)
	})
}

func (implicitFractal *ImplicitFractal2) Get4D(x, y, z, w float64) float64 {
	return implicitFractal.execute(func(frequency float64) float64 {
		return implicitFractal.noise.Eval4(x*frequency, y*frequency, z*frequency, w*frequency)
	})
}
