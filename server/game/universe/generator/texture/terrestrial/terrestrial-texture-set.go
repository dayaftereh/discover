package terrestrial

import (
	"fmt"
	"log"

	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
)

var (
	Black = textureset.NewRGBColor(0.0, 0.0, 0.0)
)

type TerrestrialTextureSet struct {
	Tiles map[int]*TerrestrialTile
}

func NewTerrestrialTextureSet() *TerrestrialTextureSet {
	return &TerrestrialTextureSet{
		Tiles: make(map[int]*TerrestrialTile),
	}
}

func (textureSet *TerrestrialTextureSet) HeightMapFrequency() float64 {
	return 1.25

}

func (textureSet *TerrestrialTextureSet) HeightMapOctaves() int64 {
	return 8
}

func (textureSet *TerrestrialTextureSet) HeatMapFrequency() float64 {
	return 3.0
}

func (textureSet *TerrestrialTextureSet) HeatMapOctaves() int64 {
	return 4
}

func (textureSet *TerrestrialTextureSet) MoistureMapFrequency() float64 {
	return 3.0
}

func (textureSet *TerrestrialTextureSet) MoistureMapOctaves() int64 {
	return 4
}

func (textureSet *TerrestrialTextureSet) findTile(tile *textureset.Tile) *TerrestrialTile {
	terrestrialTile, ok := textureSet.Tiles[tile.Index]
	if !ok {
		terrestrialTile = &TerrestrialTile{}
		textureSet.Tiles[tile.Index] = terrestrialTile
	}
	return terrestrialTile
}

func (terrestrialSet *TerrestrialTextureSet) Init(tile *textureset.Tile) {
	terrestrialTile := terrestrialSet.findTile(tile)

	// find the HeightType
	terrestrialTile.HeightType = FindHeightType(tile.HeightValue)

	fmt.Println(tile.HeightValue, terrestrialTile.HeightType)

	// get the heat value
	heatValue := tile.HeatValue

	// Adjust Heat Map based on Height - Higher == colder
	if terrestrialTile.HeightType == Forest { // Forest
		heatValue -= 0.01 * tile.HeightValue
	} else if terrestrialTile.HeightType == Rock { // Rock
		heatValue -= 0.025 * tile.HeightValue
	} else if terrestrialTile.HeightType == Snow { // Snow
		heatValue -= 0.04 * tile.HeightValue
	} else {
		heatValue += 0.001 * tile.HeightValue
	}

	// find the HeatType
	terrestrialTile.HeatType = FindHeatType(heatValue)

	// get the moisture value
	moistureValue := tile.MoistureValue

	//adjust moisture based on height
	if terrestrialTile.HeightType == DeepWater { //DeepWater
		moistureValue += 0.08 * tile.HeightValue
	} else if terrestrialTile.HeightType == ShallowWater { //ShallowWater
		moistureValue += 0.03 * tile.HeightValue
	} else if terrestrialTile.HeightType == Shore { //Shore
		moistureValue += 0.02 * tile.HeightValue
	} else if terrestrialTile.HeightType == Sand { //Sand
		moistureValue += 0.02 * tile.HeightValue
	}

	// find the moistureType
	terrestrialTile.MoistureType = FindMoistureType(moistureValue)

	// find the BiomeType
	terrestrialTile.BiomeType = FindBiomeType(terrestrialTile.MoistureType, terrestrialTile.HeatType)

}

func (terrestrialSet *TerrestrialTextureSet) HeightColor(tile *textureset.Tile) *textureset.Color {
	// find the terrestrialTile
	terrestrialTile := terrestrialSet.findTile(tile)
	// try to find the color for the type
	color, ok := HeightMapColors[terrestrialTile.HeightType]

	log.Println(color)

	if !ok {
		return Black
	}
	return color
}

func (terrestrialSet *TerrestrialTextureSet) bitmask(tile *textureset.Tile) int {
	terrestrialTile := terrestrialSet.findTile(tile)

	terrestrialTileTop := terrestrialSet.findTile(tile.Top)
	terrestrialTileBottom := terrestrialSet.findTile(tile.Bottom)
	terrestrialTileLeft := terrestrialSet.findTile(tile.Left)
	terrestrialTileRight := terrestrialSet.findTile(tile.Right)

	count := 0
	if terrestrialTile.HeightType == terrestrialTileTop.HeightType {
		count++
	} else if terrestrialTile.HeightType == terrestrialTileRight.HeightType {
		count += 2
	} else if terrestrialTile.HeightType == terrestrialTileBottom.HeightType {
		count += 4
	} else if terrestrialTile.HeightType == terrestrialTileLeft.HeightType {
		count += 8
	}

	return count

}

func (terrestrialSet *TerrestrialTextureSet) HeatColor(tile *textureset.Tile) *textureset.Color {
	// find the terrestrialTile
	terrestrialTile := terrestrialSet.findTile(tile)

	// try to find the color for the type
	color, ok := HeatMapColors[terrestrialTile.HeatType]

	if !ok {
		return Black
	}

	// do not dark edge tile for DeepWater and ShallowWater
	if terrestrialTile.HeightType == DeepWater || terrestrialTile.HeightType == ShallowWater {
		return color
	}

	bitmask := terrestrialSet.bitmask(tile)

	if bitmask != 15 {
		color = color.Lerp(Black, 0.4)
	}

	return color
}

func (terrestrialSet *TerrestrialTextureSet) MoistureColor(tile *textureset.Tile) *textureset.Color {
	// find the terrestrialTile
	terrestrialTile := terrestrialSet.findTile(tile)

	// try to find the color for the type
	color, ok := MoistureMapColors[terrestrialTile.MoistureType]

	if !ok {
		return MoistureMapColors[Wettest]
	}

	return color

}

func (terrestrialSet *TerrestrialTextureSet) biomeBitmask(tile *textureset.Tile) int {
	terrestrialTile := terrestrialSet.findTile(tile)

	terrestrialTileTop := terrestrialSet.findTile(tile.Top)
	terrestrialTileBottom := terrestrialSet.findTile(tile.Bottom)
	terrestrialTileLeft := terrestrialSet.findTile(tile.Left)
	terrestrialTileRight := terrestrialSet.findTile(tile.Right)

	count := 0
	if terrestrialTile.BiomeType == terrestrialTileTop.BiomeType {
		count++
	} else if terrestrialTile.BiomeType == terrestrialTileRight.BiomeType {
		count += 2
	} else if terrestrialTile.BiomeType == terrestrialTileBottom.BiomeType {
		count += 4
	} else if terrestrialTile.BiomeType == terrestrialTileLeft.BiomeType {
		count += 8
	}

	return count
}

func (terrestrialSet *TerrestrialTextureSet) BiomeColor(tile *textureset.Tile) *textureset.Color {
	// find the terrestrialTile
	terrestrialTile := terrestrialSet.findTile(tile)

	// Water tiles
	if terrestrialTile.HeightType == DeepWater {
		return HeightMapColors[DeepWater]
	} else if terrestrialTile.HeightType == ShallowWater {
		return HeightMapColors[ShallowWater]
	}

	// try to find the color for the type
	color, ok := BiomeMapColors[terrestrialTile.BiomeType]

	if !ok {
		return Black
	}

	// calculate the biomeBitmask
	biomeBitmask := terrestrialSet.biomeBitmask(tile)

	if biomeBitmask != 15 {
		color = color.Lerp(Black, 0.4)
	}

	return color
}

func (textureSet *TerrestrialTextureSet) Cloud1Color(tile *textureset.Tile) *textureset.Color {
	return nil
}

func (textureSet *TerrestrialTextureSet) Cloud2Color(tile *textureset.Tile) *textureset.Color {
	return nil
}
