package textureset

type TextureSet interface {
	// noise
	HeightMapFrequency() float64
	HeightMapOctaves() int64
	HeatMapFrequency() float64
	HeatMapOctaves() int64
	MoistureMapFrequency() float64
	MoistureMapOctaves() int64

	// tiles
	Init(tile *Tile)
	HeightColor(tile *Tile) *Color
	HeatColor(tile *Tile) *Color
	MoistureColor(tile *Tile) *Color
	BiomeColor(tile *Tile) *Color
	Cloud1Color(tile *Tile) *Color
	Cloud2Color(tile *Tile) *Color
}
