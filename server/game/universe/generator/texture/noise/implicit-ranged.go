package noise

type ImplicitRanged struct {
	source ImplicitBase
	min    float64
	max    float64
}

func NewImplicitRanged(source ImplicitBase, min, max float64) *ImplicitRanged {
	return &ImplicitRanged{
		min:    min,
		max:    max,
		source: source,
	}
}

func (implicitRanged *ImplicitRanged) execute(fn func() float64) float64 {
	value := fn()
	return (value * (implicitRanged.max - implicitRanged.min)) + implicitRanged.min
}

func (implicitRanged *ImplicitRanged) Get2D(x, y float64) float64 {
	return implicitRanged.execute(func() float64 {
		return implicitRanged.source.Get2D(x, y)
	})
}

func (implicitRanged *ImplicitRanged) Get3D(x, y, z float64) float64 {
	return implicitRanged.execute(func() float64 {
		return implicitRanged.source.Get3D(x, y, z)
	})
}

func (implicitRanged *ImplicitRanged) Get4D(x, y, z, w float64) float64 {
	return implicitRanged.execute(func() float64 {
		return implicitRanged.source.Get4D(x, y, z, w)
	})
}
