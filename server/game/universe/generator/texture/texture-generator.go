package texture

import (
	"image"
	"image/color"
	"math"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
	textureset "github.com/dayaftereh/discover/server/game/universe/generator/texture/texture-set"
)

type TextureGenerator struct {
	seed   int64
	width  int
	height int
	// Tiles
	tiles []*textureset.Tile
	// texture set
	textureSet textureset.TextureSet
}

func NewTextureGenerator(width, height int, textureSet textureset.TextureSet, seed int64) *TextureGenerator {
	return &TextureGenerator{
		width:      width,
		height:     height,
		seed:       seed,
		textureSet: textureSet,
		tiles:      make([]*textureset.Tile, width*height),
	}
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
	heightMapFrequency := textureGenerator.textureSet.HeightMapFrequency()
	heightMapOctaves := textureGenerator.textureSet.HeightMapOctaves()

	heightMap := noise.NewImplicitFractal(heightMapFrequency, heightMapOctaves, true, textureGenerator.seed)

	// heatMap
	heatMapFrequency := textureGenerator.textureSet.HeatMapFrequency()
	heatMapOctaves := textureGenerator.textureSet.HeatMapOctaves()

	heatFractal := noise.NewImplicitFractal(heatMapFrequency, heatMapOctaves, true, textureGenerator.seed)
	gradient := noise.NewImplicitGradient(1.0, 1.0, 0.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0)
	heatMap := noise.NewImplicitCombiner(noise.CombinerTypeMultiply)
	heatMap.AddSource(gradient)
	heatMap.AddSource(heatFractal)

	// MoistureMap
	moistureMapFrequency := textureGenerator.textureSet.MoistureMapFrequency()
	moistureMapOctaves := textureGenerator.textureSet.MoistureMapOctaves()

	moistureMap := noise.NewImplicitFractal(moistureMapFrequency, moistureMapOctaves, true, textureGenerator.seed)

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

		// get the tile
		tile := textureGenerator.getTile(x, y)

		tile.HeightValue = heightMap.Get4D(nx, ny, nz, nw)
		tile.HeatValue = heatMap.Get4D(nx, ny, nz, nw)
		tile.MoistureValue = moistureMap.Get4D(nx, ny, nz, nw)

	})

	textureGenerator.forEachPixel(func(x, y int) {
		// get the tile
		tile := textureGenerator.getTile(x, y)
		// initialize the tile
		tile.Init()
	})
}

func (textureGenerator *TextureGenerator) tileIndex(x, y int) int {
	x = x % textureGenerator.width
	y = y % textureGenerator.height

	if x < 0 {
		x = textureGenerator.width + x
	}

	if y < 0 {
		y = textureGenerator.height + y
	}

	index := y*textureGenerator.width + x
	return index
}

func (textureGenerator *TextureGenerator) getTile(x, y int) *textureset.Tile {
	index := textureGenerator.tileIndex(x, y)
	tile := textureGenerator.tiles[index]

	if tile == nil {
		// create the tile
		tile = &textureset.Tile{
			Index:      index,
			X:          x,
			Y:          y,
			TextureSet: textureGenerator.textureSet,
		}
		// store the tile
		textureGenerator.tiles[index] = tile

		// locate the tile
		tile.Top = textureGenerator.getTile(x, y-1)
		tile.Bottom = textureGenerator.getTile(x, y+1)
		tile.Left = textureGenerator.getTile(x-1, y)
		tile.Right = textureGenerator.getTile(x+1, y)
	}

	return tile
}

func (textureGenerator *TextureGenerator) imageFromEachTile(fn func(tile *textureset.Tile) *textureset.Color) *image.RGBA {
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
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.HeightColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateHeatMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.HeatColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateMoistureMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.MoistureColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateBiomeMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.BiomeColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) color2RGBA(c *textureset.Color) color.RGBA {
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
