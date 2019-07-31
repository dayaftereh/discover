package noise

type ImplicitBase interface {
	Get2D(x, y float64) float64
	Get3D(x, y, z float64) float64
	Get4D(x, y, z, w float64) float64
}
