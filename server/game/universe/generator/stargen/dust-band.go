package stargen

import "fmt"

type DustBand struct {
	InnerEdge   float64
	OuterEdge   float64
	DustPresent bool
	GasPresent  bool
}

func (dustBand *DustBand) String() string {
	s := fmt.Sprintf("DustBand: [\n")
	s = fmt.Sprintf("%s InnerEdge: %f\n", s, dustBand.InnerEdge)
	s = fmt.Sprintf("%s OuterEdge: %f\n", s, dustBand.OuterEdge)
	s = fmt.Sprintf("%s DustPresent: %v\n", s, dustBand.DustPresent)
	s = fmt.Sprintf("%s GasPresent: %v\n", s, dustBand.GasPresent)
	s = fmt.Sprintf("%s]\n", s)
	return s
}
