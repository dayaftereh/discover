package textureset

type Tile struct {
	Index         int
	X             int
	Y             int
	HeightValue   float64
	HeatValue     float64
	MoistureValue float64
	Top           *Tile
	Bottom        *Tile
	Left          *Tile
	Right         *Tile
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

func (tile *Tile) Cloud1Color() *Color {
	return tile.TextureSet.Cloud1Color(tile)
}

func (tile *Tile) Cloud2Color() *Color {
	return tile.TextureSet.Cloud2Color(tile)
}
