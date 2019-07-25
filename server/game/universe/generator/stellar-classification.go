package generator

import (
	"math"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/mathf"
)

type StellarClass struct {
	Class       string
	Color       int64
	Mass        *mathf.Range
	Radius      *mathf.Range
	Temperature *mathf.Range
	Luminosity  *mathf.Range
}

// Stellar classification https://en.wikipedia.org/wiki/Stellar_classification

var StellarClassification []*StellarClass = []*StellarClass{
	&StellarClass{
		Class:       "O",
		Color:       255,                                // blue
		Mass:        mathf.NewRange(16.0, math.Inf(1)),  // >= 16 M
		Radius:      mathf.NewRange(6.6, math.Inf(1)),   // >= 6.6 R
		Temperature: mathf.NewRange(30, math.Inf(1)),    // >= 30 *1000 K
		Luminosity:  mathf.NewRange(30000, math.Inf(1)), // >= 30000 L
	},
	&StellarClass{
		Class:       "B",
		Color:       10079487,                  // blue white
		Mass:        mathf.NewRange(2.1, 16.0), // 2.1–16 M
		Radius:      mathf.NewRange(1.8, 6.6),  // 1.8–6.6 R
		Temperature: mathf.NewRange(10, 30),    // 10 - 30 *1000 K
		Luminosity:  mathf.NewRange(25, 30000), // 25 - 30000 L
	},
	&StellarClass{
		Class:       "A",
		Color:       16777215,                 // white
		Mass:        mathf.NewRange(1.4, 2.1), // 1.4 - 2.1 M
		Radius:      mathf.NewRange(1.4, 1.8), // 1.4 - 1.8 R
		Temperature: mathf.NewRange(7.5, 10),  // 7.5 - 10 *1000 K
		Luminosity:  mathf.NewRange(5, 25),    // 5 - 25 L
	},
	&StellarClass{
		Class:       "F",
		Color:       16777164,                  // yellow white
		Mass:        mathf.NewRange(1.04, 1.4), // 1.04 - 1.4 M
		Radius:      mathf.NewRange(1.15, 1.4), // 1.15 - 1.4 R
		Temperature: mathf.NewRange(6, 7.5),    // 6 - 7.5 *1000 K
		Luminosity:  mathf.NewRange(1.5, 5),    // 1.5 - 5 L
	},
	&StellarClass{
		Class:       "G",
		Color:       16776960,                   // yellow
		Mass:        mathf.NewRange(0.8, 1.04),  // 0.8 - 1.04 M
		Radius:      mathf.NewRange(0.96, 1.15), // 0.96 - 1.15 R
		Temperature: mathf.NewRange(5.2, 6.0),   // 5.2 - 6 *1000 K
		Luminosity:  mathf.NewRange(0.6, 1.5),   // 0.6 - 1.5 L
	},
	&StellarClass{
		Class:       "K",
		Color:       16761446,                  // light orange
		Mass:        mathf.NewRange(0.45, 0.8), // 0.45 - 0.8 M
		Radius:      mathf.NewRange(0.7, 0.96), // 0.7 - 0.96 R
		Temperature: mathf.NewRange(3.7, 5.2),  // 3.7 - 5.2 *1000 K
		Luminosity:  mathf.NewRange(0.08, 0.6), // 0.08 - 0.6 L
	},
	&StellarClass{
		Class:       "M",
		Color:       13395456,                   // orange red
		Mass:        mathf.NewRange(0.08, 0.45), // 0.08 - 0.45 M
		Radius:      mathf.NewRange(0.0, 0.7),   // 0 - 0.7 R
		Temperature: mathf.NewRange(2.4, 3.7),   // 2.4 - 3.7 *1000 K
		Luminosity:  mathf.NewRange(0.0, 0.08),  // 0.0 - 0.08 L
	},
}

func RandStellarClass() *StellarClass {
	length := len(StellarClassification)
	index := utils.RandIntn(length)
	class := StellarClassification[index]
	return class
}
