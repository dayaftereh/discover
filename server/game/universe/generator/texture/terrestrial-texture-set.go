package texture

var (
	// Height Map Colors
	TerrestrialDeepColor = NewRGBColor255(15.0, 30.0, 80.0)
)

type TerrestrialTextureSet struct {
}

func (textureSet *TerrestrialTextureSet) HeightColor(tile *Tile) *Color {
	return nil
}
func (textureSet *TerrestrialTextureSet) HeatColor(tile *Tile) *Color {
	return nil
}

func (textureSet *TerrestrialTextureSet) MoistureColor(tile *Tile) *Color {
	return nil
}

func (textureSet *TerrestrialTextureSet) BiomeColor(tile *Tile) *Color {
	return nil
}

func (textureSet *TerrestrialTextureSet) Cloud1Color(tile *Tile) *Color {
	return nil
}

func (textureSet *TerrestrialTextureSet) Cloud2Color(tile *Tile) *Color {
	return nil
}
