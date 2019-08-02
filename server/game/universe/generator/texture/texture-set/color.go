package textureset

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

func (color *Color) Lerp(other *Color, x float64) *Color {
	return NewRGBAColor(
		x*(other.R-color.R)+color.R,
		x*(other.G-color.G)+color.G,
		x*(other.B-color.B)+color.B,
		x*(other.A-color.A)+color.A,
	)
}

func (color *Color) Grayscale() float64 {
	return color.R*0.299 + color.G*0.587 + color.B*0.114
}
