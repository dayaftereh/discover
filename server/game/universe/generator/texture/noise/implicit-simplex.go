package noise

import "github.com/ojrac/opensimplex-go"

type ImplicitSimplex struct {
	noise opensimplex.Noise
}

func NewImplicitSimplex(seed int64) *ImplicitSimplex {
	return &ImplicitSimplex{
		noise: opensimplex.NewNormalized(seed),
	}
}

func (implicitSimplex *ImplicitSimplex) Get2D(x, y float64) float64 {
	return implicitSimplex.noise.Eval2(x, y)
}

func (implicitSimplex *ImplicitSimplex) Get3D(x, y, z float64) float64 {
	return implicitSimplex.noise.Eval3(x, y, z)
}

func (implicitSimplex *ImplicitSimplex) Get4D(x, y, z, w float64) float64 {
	return implicitSimplex.noise.Eval4(x, y, z, w)
}
