package texture

type TextureSet interface {
	HeightColor(tile *Tile) *Color
	HeatColor(tile *Tile) *Color
	MoistureColor(tile *Tile) *Color
	BiomeColor(tile *Tile) *Color
	Cloud1Color(tile *Tile) *Color
	Cloud2Color(tile *Tile) *Color
}
