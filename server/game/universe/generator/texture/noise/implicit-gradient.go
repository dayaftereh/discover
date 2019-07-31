package noise

type ImplicitGradient struct {
	gradientX0 float64
	gradientY0 float64
	gradientZ0 float64
	gradientW0 float64
	gradientU0 float64
	gradientV0 float64
	gradientX1 float64
	gradientY1 float64
	gradientZ1 float64
	gradientW1 float64
	gradientU1 float64
	gradientV1 float64
	length2    float64
	length3    float64
	length4    float64
	length6    float64
}

func NewImplicitGradient(x0, x1, y0, y1, z0, z1, w0, w1, u0, u1, v0, v1 float64) *ImplicitGradient {
	i := &ImplicitGradient{
		gradientX0: x0,
		gradientY0: y0,
		gradientZ0: z0,
		gradientW0: w0,
		gradientU0: u0,
		gradientV0: v0,

		gradientX1: x1 - x0,
		gradientY1: y1 - y0,
		gradientZ1: z1 - z0,
		gradientW1: w1 - w0,
		gradientU1: u1 - u0,
		gradientV1: v1 - v0,
	}

	i.length2 = (i.gradientX1*i.gradientX1 + i.gradientY1*i.gradientY1)
	i.length3 = (i.gradientX1*i.gradientX1 + i.gradientY1*i.gradientY1 + i.gradientZ1*i.gradientZ1)
	i.length4 = (i.gradientX1*i.gradientX1 + i.gradientY1*i.gradientY1 + i.gradientZ1*i.gradientZ1 + i.gradientW1*i.gradientW1)
	i.length6 = (i.gradientX1*i.gradientX1 + i.gradientY1*i.gradientY1 + i.gradientZ1*i.gradientZ1 + i.gradientW1*i.gradientW1 + i.gradientU1*i.gradientU1 + i.gradientV1*i.gradientV1)

	return i
}

func (implicitGradient *ImplicitGradient) Get2D(x float64, y float64) float64 {
	dx := x - implicitGradient.gradientX0
	dy := y - implicitGradient.gradientY0
	dp := dx*implicitGradient.gradientX1 + dy*implicitGradient.gradientY1
	dp /= implicitGradient.length2
	return dp
}

func (implicitGradient *ImplicitGradient) Get3D(x float64, y float64, z float64) float64 {
	dx := x - implicitGradient.gradientX0
	dy := y - implicitGradient.gradientY0
	dz := z - implicitGradient.gradientZ0
	dp := dx*implicitGradient.gradientX1 + dy*implicitGradient.gradientY1 + dz*implicitGradient.gradientZ1
	dp /= implicitGradient.length3
	return dp
}

func (implicitGradient *ImplicitGradient) Get4D(x float64, y float64, z float64, w float64) float64 {
	dx := x - implicitGradient.gradientX0
	dy := y - implicitGradient.gradientY0
	dz := z - implicitGradient.gradientZ0
	dw := w - implicitGradient.gradientW0
	dp := dx*implicitGradient.gradientX1 + dy*implicitGradient.gradientY1 + dz*implicitGradient.gradientZ1 + dw*implicitGradient.gradientW1
	dp /= implicitGradient.length4
	return dp
}
