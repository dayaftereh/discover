package texture

import (
	"image"
	"image/color"
	"math"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
)

type TextureGenerator struct {
	seed   int64
	width  int
	height int
	//  heightMap
	heightMapFrequency float64
	heightMapOctaves   int64
	//  heatMap
	heatMapFrequency float64
	heatMapOctaves   int64
	// moistureMap
	moistureMapFrequency float64
	moistureMapOctaves   int64
	// Tiles
	tiles []*Tile
	// texture set
	textureSet TextureSet
}

func (textureGenerator *TextureGenerator) forEachPixel(fn func(x, y int)) {
	for x := 0; x < textureGenerator.width; x++ {
		for y := 0; y < textureGenerator.height; y++ {
			fn(x, y)
		}
	}
}

func (textureGenerator *TextureGenerator) Init() {
	// heightMap
	heightMap := noise.NewImplicitFractal(textureGenerator.heightMapFrequency, textureGenerator.heightMapOctaves, true, textureGenerator.seed)
	// heatMap
	heatFractal := noise.NewImplicitFractal(textureGenerator.heatMapFrequency, textureGenerator.heatMapOctaves, true, textureGenerator.seed)
	gradient := noise.NewImplicitGradient(1.0, 1.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0)
	heatMap := noise.NewImplicitCombiner(noise.CombinerTypeMultiply)
	heatMap.AddSource(gradient)
	heatMap.AddSource(heatFractal)
	// MoistureMap
	moistureMap := noise.NewImplicitFractal(textureGenerator.moistureMapFrequency, textureGenerator.moistureMapOctaves, true, textureGenerator.seed)

	// initialize the map data
	heightMapData := NewMapData(textureGenerator.width, textureGenerator.height)
	heatMapData := NewMapData(textureGenerator.width, textureGenerator.height)
	moistureMapData := NewMapData(textureGenerator.width, textureGenerator.height)

	// generate all map data
	textureGenerator.forEachPixel(func(x, y int) {
		// WRAP ON BOTH AXIS

		// Noise range
		x1, x2 := 0.0, 2.0
		y1, y2 := 0.0, 2.0
		dx := x2 - x1
		dy := y2 - y1

		// Sample noise at smaller intervals
		s := float64(x) / float64(textureGenerator.width)
		t := float64(y) / float64(textureGenerator.height)

		// Calculate our 4D coordinates
		nx := x1 + math.Cos(s*2.0*math.Pi)*dx/(2.0*math.Pi)
		ny := y1 + math.Cos(t*2.0*math.Pi)*dy/(2.0*math.Pi)
		nz := x1 + math.Sin(s*2.0*math.Pi)*dx/(2.0*math.Pi)
		nw := y1 + math.Sin(t*2.0*math.Pi)*dy/(2.0*math.Pi)

		heightValue := heightMap.Get4D(nx, ny, nz, nw)
		heatValue := heatMap.Get4D(nx, ny, nz, nw)
		moistureValue := moistureMap.Get4D(nx, ny, nz, nw)

		heightMapData.Set(x, y, heightValue)
		heatMapData.Set(x, y, heatValue)
		moistureMapData.Set(x, y, moistureValue)
	})

	textureGenerator.forEachPixel(func(x, y int) {
		tile := textureGenerator.getTile(x, y)
		tile.HeightValue = (heightMapData.Get(x, y) - heatMapData.MinValue) / (heatMapData.MaxValue - heatMapData.MinValue)
		tile.HeatValue = (heatMapData.Get(x, y) - heatMapData.MinValue) / (heatMapData.MaxValue - heatMapData.MinValue)
		tile.MoistureValue = (moistureMapData.Get(x, y) - moistureMapData.MinValue) / (moistureMapData.MaxValue - moistureMapData.MinValue)
	})

}

func (textureGenerator *TextureGenerator) tileIndex(x, y int) int {
	x = int(math.Mod(float64(x), float64(textureGenerator.width)))
	y = int(math.Mod(float64(y), float64(textureGenerator.height)))
	index := y*textureGenerator.width + x
	return index
}

func (textureGenerator *TextureGenerator) getTile(x, y int) *Tile {
	index := textureGenerator.tileIndex(x, y)
	tile := textureGenerator.tiles[index]

	if tile == nil {
		tile = &Tile{
			X:          x,
			Y:          y,
			TextureSet: textureGenerator.textureSet,
			Top:        textureGenerator.getTile(x, y-1),
			Bottom:     textureGenerator.getTile(x, y+1),
			Left:       textureGenerator.getTile(x-1, y),
			Right:      textureGenerator.getTile(x+1, y),
		}
		textureGenerator.tiles[index] = tile
	}
	return tile
}

func (textureGenerator *TextureGenerator) imageFromEachTile(fn func(tile *Tile) *Color) *image.RGBA {
	rect := image.Rect(0, 0, textureGenerator.width, textureGenerator.height)
	img := image.NewRGBA(rect)

	textureGenerator.forEachPixel(func(x, y int) {
		tile := textureGenerator.getTile(x, y)
		c := fn(tile)
		rgba := textureGenerator.color2RGBA(c)
		img.Set(x, y, rgba)
	})

	return img
}

func (textureGenerator *TextureGenerator) GenerateHeightMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *Tile) *Color {
		return tile.HeightColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateHeatMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *Tile) *Color {
		return tile.HeatColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateMoistureMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *Tile) *Color {
		return tile.MoistureColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateBiomeMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *Tile) *Color {
		return tile.MoistureColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) color2RGBA(c *Color) color.RGBA {
	r := uint8(255 * mathf.Clamp(c.R, 0.0, 1.0))
	g := uint8(255 * mathf.Clamp(c.G, 0.0, 1.0))
	b := uint8(255 * mathf.Clamp(c.B, 0.0, 1.0))
	a := uint8(255 * mathf.Clamp(c.A, 0.0, 1.0))

	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
