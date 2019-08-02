package textureset

type Tile struct {
	Index         int
	X             int
	Y             int
	HeightValue   float64
	HeatValue     float64
	MoistureValue float64
	CloudValue    float64
	Top           func() *Tile
	Bottom        func() *Tile
	Left          func() *Tile
	Right         func() *Tile
	TextureSet    TextureSet
}

func (tile *Tile) Init() {
	tile.TextureSet.Init(tile)
}

func (tile *Tile) HeightColor() *Color {
	return tile.TextureSet.HeightColor(tile)
}

func (tile *Tile) HeatColor() *Color {
	return tile.TextureSet.HeatColor(tile)
}

func (tile *Tile) MoistureColor() *Color {
	return tile.TextureSet.MoistureColor(tile)
}

func (tile *Tile) BiomeColor() *Color {
	return tile.TextureSet.BiomeColor(tile)
}

func (tile *Tile) BumpColor() *Color {
	return tile.TextureSet.BumpColor(tile)
}

func (tile *Tile) CloudColor() *Color {
	return tile.TextureSet.CloudColor(tile)
}

func (tile *Tile) SpecularColor() *Color {
	return tile.TextureSet.SpecularColor(tile)
}
