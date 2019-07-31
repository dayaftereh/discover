package texture

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func NewRGBColor(r, g, b float64) *Color {
	return &Color{
		R: r,
		G: g,
		B: b,
		A: 1.0,
	}
}

func NewRGBColor255(r, g, b float64) *Color {
	return &Color{
		R: r / 255.0,
		G: g / 255.0,
		B: b / 255.0,
		A: 1.0,
	}
}

func NewRGBAColor(r, g, b, a float64) *Color {
	return &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
