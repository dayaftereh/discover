package noise

type ImplicitInputTransform struct {
	source    ImplicitBase
	transform func(value float64) float64
}

func NewImplicitInputTransform(source ImplicitBase, transform func(value float64) float64) *ImplicitInputTransform {
	return &ImplicitInputTransform{
		source:    source,
		transform: transform,
	}
}

func (implicitInputTransform *ImplicitInputTransform) Get2D(x, y float64) float64 {
	tx := implicitInputTransform.transform(x)
	ty := implicitInputTransform.transform(y)
	return implicitInputTransform.source.Get2D(tx, ty)
}

func (implicitInputTransform *ImplicitInputTransform) Get3D(x, y, z float64) float64 {
	tx := implicitInputTransform.transform(x)
	ty := implicitInputTransform.transform(y)
	tz := implicitInputTransform.transform(z)
	return implicitInputTransform.source.Get3D(tx, ty, tz)
}

func (implicitInputTransform *ImplicitInputTransform) Get4D(x, y, z, w float64) float64 {
	tx := implicitInputTransform.transform(x)
	ty := implicitInputTransform.transform(y)
	tz := implicitInputTransform.transform(z)
	tw := implicitInputTransform.transform(w)
	return implicitInputTransform.source.Get4D(tx, ty, tz, tw)
}
