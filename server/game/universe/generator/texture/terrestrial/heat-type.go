package terrestrial

import (
	"math"

	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
)

type HeatType struct {
	Value    float64
	Range    bool
	MaxValue float64
}

// HeatTypes
var (
	HeatTypes = []HeatType{
		Coldest, Colder, Cold, Warm, Warmer, Warmest,
	}

	Coldest = HeatType{
		Value: 0.05,
	}
	Colder = HeatType{
		Value: 0.18,
	}
	Cold = HeatType{
		Value: 0.4,
	}
	Warm = HeatType{
		Value: 0.6,
	}
	Warmer = HeatType{
		Value: 0.8,
	}
	Warmest = HeatType{
		Value:    0.8,
		Range:    true,
		MaxValue: math.Inf(1),
	}

	// Heat Colors
	HeatMapColors = map[HeatType]*textureset.Color{
		Coldest: textureset.NewRGBColor(0.0, 1.0, 1.0),
		Colder:  textureset.NewRGBColor255(170, 255, 255),
		Cold:    textureset.NewRGBColor255(0.0, 229.0, 133.0),
		Warm:    textureset.NewRGBColor255(255.0, 255.0, 100.0),
		Warmer:  textureset.NewRGBColor255(255.0, 100.0, 0.0),
		Warmest: textureset.NewRGBColor255(241.0, 12.0, 0.0),
	}
)

func FindHeatType(value float64) HeatType {
	for _, typ := range HeatTypes {
		if typ.Range {
			if typ.Value <= value && value <= typ.MaxValue {
				return typ
			}
		}
		if value < typ.Value {
			return typ
		}
	}

	return Warmest
}
