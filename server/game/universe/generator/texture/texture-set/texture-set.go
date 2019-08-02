package textureset

type TextureSet interface {
	// noise
	HeightMapFrequency() float64
	HeightMapOctaves() int64
	HeatMapFrequency() float64
	HeatMapOctaves() int64
	MoistureMapFrequency() float64
	MoistureMapOctaves() int64
	CloudMapFrequency() float64
	CloudMapOctaves() int64

	// tiles
	Init(tile *Tile)
	HeightColor(tile *Tile) *Color
	HeatColor(tile *Tile) *Color
	MoistureColor(tile *Tile) *Color
	BiomeColor(tile *Tile) *Color
	BumpColor(tile *Tile) *Color
	CloudColor(tile *Tile) *Color
	SpecularColor(tile *Tile) *Color
}
