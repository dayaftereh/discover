package mathf

type Range struct {
	Min float64
	Max float64
}

func NewRange(min float64, max float64) *Range {
	return &Range{
		Min: min,
		Max: max,
	}
}
