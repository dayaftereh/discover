package terrestrial

import (
	"math"

	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
)

type HeightType struct {
	Value    float64
	Range    bool
	MaxValue float64
}

var (
	HeightTypes = []*HeightType{
		DeepWater, ShallowWater, Shore, Sand, Grass, Forest, Rock, Snow,
	}

	DeepWater = &HeightType{
		Value: 0.2,
	}
	ShallowWater = &HeightType{
		Value: 0.4,
	}
	Shore = &HeightType{
		Value: 0.5,
	}
	Sand = &HeightType{
		Value: 0.55,
	}
	Grass = &HeightType{
		Value: 0.7,
	}
	Forest = &HeightType{
		Value: 0.8,
	}
	Rock = &HeightType{
		Value: 0.9,
	}
	Snow = &HeightType{
		Value:    0.9,
		Range:    true,
		MaxValue: math.Inf(1),
	}
	River = &HeightType{
		Value: 0.2,
	}

	// Height Colors
	HeightMapColors = map[*HeightType]*textureset.Color{
		DeepWater:    textureset.NewRGBColor255(15.0, 30.0, 80.0),
		ShallowWater: textureset.NewRGBColor255(15.0, 40.0, 90.0),
		Shore:        textureset.NewRGBColor255(198.0, 190.0, 31.0),
		Sand:         textureset.NewRGBColor255(198.0, 190.0, 31.0),
		Grass:        textureset.NewRGBColor255(50.0, 220.0, 20.0),
		Forest:       textureset.NewRGBColor255(16.0, 160.0, 0.0),
		Rock:         textureset.NewRGBColor(0.5, 0.5, 0.5),
		Snow:         textureset.NewRGBColor(1.0, 1.0, 1.0),
		River:        textureset.NewRGBColor255(241.0, 12.0, 0.0),
	}

	BumpMapColors = map[*HeightType]*textureset.Color{
		DeepWater:    textureset.NewRGBColor(0.0, 0.0, 0.0),
		ShallowWater: textureset.NewRGBColor(0.0, 0.0, 0.0),
		Shore:        textureset.NewRGBColor(0.3, 0.3, 0.3),
		Sand:         textureset.NewRGBColor(0.3, 0.3, 0.3),
		Grass:        textureset.NewRGBColor(0.45, 0.45, 0.45),
		Forest:       textureset.NewRGBColor(0.6, 0.6, 0.6),
		Rock:         textureset.NewRGBColor(0.75, 0.75, 0.75),
		Snow:         textureset.NewRGBColor(1.0, 1.0, 1.0),
		River:        textureset.NewRGBColor255(0.0, 0.0, 0.0),
	}
)

func FindHeightType(value float64) *HeightType {
	for _, typ := range HeightTypes {
		if typ.Range {
			if typ.Value <= value && value <= typ.MaxValue {
				return typ
			}
		}
		if value < typ.Value {
			return typ
		}
	}

	return Snow
}
