package terrestrial

import (
	"math"

	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
)

type MoistureType struct {
	Value    float64
	Range    bool
	MaxValue float64
}

var (
	MoistureTypes = []*MoistureType{
		Dryest, Dryer, Dry, Wet, Wetter, Wettest,
	}

	Wettest = &MoistureType{
		Value:    0.95,
		Range:    true,
		MaxValue: math.Inf(1),
	}
	Wetter = &MoistureType{
		Value: 0.95,
	}
	Wet = &MoistureType{
		Value: 0.8,
	}
	Dry = &MoistureType{
		Value: 0.6,
	}
	Dryer = &MoistureType{
		Value: 0.4,
	}
	Dryest = &MoistureType{
		Value: 0.27,
	}

	// Moisture Colors

	MoistureMapColors = map[*MoistureType]*textureset.Color{
		Wettest: textureset.NewRGBColor255(0.0, 0.0, 100.0),
		Wetter:  textureset.NewRGBColor255(20.0, 70.0, 255.0),
		Wet:     textureset.NewRGBColor255(85.0, 255.0, 255.0),
		Dry:     textureset.NewRGBColor255(80.0, 255.0, 0.0),
		Dryer:   textureset.NewRGBColor255(245.0, 245.0, 23.0),
		Dryest:  textureset.NewRGBColor255(255.0, 139.0, 17.0),
	}
)

func FindMoistureType(value float64) *MoistureType {
	for _, typ := range MoistureTypes {
		if typ.Range {
			if typ.Value <= value && value <= typ.MaxValue {
				return typ
			}
		}
		if value < typ.Value {
			return typ
		}
	}

	return Wettest
}
