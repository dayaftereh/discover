package terrestrial

import textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"

type BiomeType string

const (
	Desert              BiomeType = "desert"
	Savanna             BiomeType = "savanna"
	TropicalRainforest  BiomeType = "tropical-rainforset"
	Grassland           BiomeType = "grassland"
	Woodland            BiomeType = "woodland"
	SeasonalForest      BiomeType = "seasonal-forest"
	TemperateRainforest BiomeType = "temperate-rainforest"
	BorealForest        BiomeType = "boreal-forest"
	Tundra              BiomeType = "tundra"
	Ice                 BiomeType = "ice"
)

var (
	BiomeMatrix = map[MoistureType]map[HeatType]BiomeType{
		Dryest: map[HeatType]BiomeType{
			Coldest: Ice, Colder: Tundra, Cold: Grassland, Warm: Desert, Warmer: Desert, Warmest: Desert,
		},
		Dryer: map[HeatType]BiomeType{
			Coldest: Ice, Colder: Tundra, Cold: Grassland, Warm: Desert, Warmer: Desert, Warmest: Desert,
		},
		Dry: map[HeatType]BiomeType{
			Coldest: Ice, Colder: Tundra, Cold: Woodland, Warm: Woodland, Warmer: Savanna, Warmest: Savanna,
		},
		Wet: map[HeatType]BiomeType{
			Coldest: Ice, Colder: Tundra, Cold: BorealForest, Warm: Woodland, Warmer: Savanna, Warmest: Savanna,
		},
		Wetter: map[HeatType]BiomeType{
			Coldest: Ice, Colder: Tundra, Cold: BorealForest, Warm: SeasonalForest, Warmer: TropicalRainforest, Warmest: TropicalRainforest,
		},
		Wettest: map[HeatType]BiomeType{
			Coldest: Ice, Colder: Tundra, Cold: BorealForest, Warm: TropicalRainforest, Warmer: TropicalRainforest, Warmest: TropicalRainforest,
		},
	}

	// Biome Colors
	BiomeMapColors = map[BiomeType]*textureset.Color{
		Desert:              textureset.NewRGBColor255(238.0, 218.0, 130.0),
		Savanna:             textureset.NewRGBColor255(177.0, 209.0, 110.0),
		TropicalRainforest:  textureset.NewRGBColor255(66.0, 123.0, 25.0),
		Grassland:           textureset.NewRGBColor255(164.0, 225.0, 99.0),
		Woodland:            textureset.NewRGBColor255(139.0, 175.0, 90.0),
		SeasonalForest:      textureset.NewRGBColor255(73.0, 100.0, 35.0),
		TemperateRainforest: textureset.NewRGBColor255(29.0, 73.0, 40.0),
		BorealForest:        textureset.NewRGBColor255(95.0, 115.0, 62.0),
		Tundra:              textureset.NewRGBColor255(96.0, 131.0, 112.0),
		Ice:                 textureset.NewRGBColor255(1.0, 1.0, 1.0),
	}
)

func FindBiomeType(moistureType MoistureType, heatType HeatType) BiomeType {
	moistureRow, ok := BiomeMatrix[moistureType]

	if !ok {
		return Ice
	}

	biomeType, ok := moistureRow[heatType]

	if !ok {
		return Ice
	}

	return biomeType
}
