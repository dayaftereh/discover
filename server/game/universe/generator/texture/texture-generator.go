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

	// CloudMap
	cloudMapFrequency := textureGenerator.textureSet.CloudMapFrequency()
	cloudMapOctaves := textureGenerator.textureSet.CloudMapOctaves()
	cloudMap := noise.NewImplicitFractal(cloudMapFrequency, cloudMapOctaves, true, textureGenerator.seed)

	// MapData
	heightData := NewMapData(textureGenerator.width, textureGenerator.height)
	heatData := NewMapData(textureGenerator.width, textureGenerator.height)
	moistureData := NewMapData(textureGenerator.width, textureGenerator.height)
	cloudData := NewMapData(textureGenerator.width, textureGenerator.height)

	southLatBound := -180.0
	northLatBound := 180.0
	westLonBound := -90.0
	eastLonBound := 90.0

	lonExtent := eastLonBound - westLonBound
	latExtent := northLatBound - southLatBound

	xDelta := lonExtent / float64(textureGenerator.width)
	yDelta := latExtent / float64(textureGenerator.height)

	curLon := westLonBound
	curLat := southLatBound

	for x := 0; x < textureGenerator.width; x++ {
		curLon = westLonBound
		for y := 0; y < textureGenerator.height; y++ {
			x1, y1, z1 := textureGenerator.latLonToXYZ(curLat, curLon)

			// set the heightValue
			heightValue := heightMap.Get3D(x1, y1, z1)
			heightData.Set(x, y, heightValue)

			// set the heatValue
			heatValue := heatMap.Get3D(x1, y1, z1)

			coldness := math.Abs(curLon) / 90.0
			heat := 1.0 - math.Abs(curLon)/90.0

			heatValue += heat
			heatValue -= coldness

			heatData.Set(x, y, heatValue)

			// set the moistureValue
			moistureValue := moistureMap.Get3D(x1, y1, z1)
			moistureData.Set(x, y, moistureValue)

			cloudValue := cloudMap.Get3D(x1, y1, z1)
			cloudData.Set(x, y, cloudValue)

			curLon += xDelta
		}
		curLat += yDelta
	}

	textureGenerator.forEachTile(func(tile *textureset.Tile) {
		tile.HeightValue = heightData.GetNormalized(tile.X, tile.Y)
		tile.HeatValue = heatData.GetNormalized(tile.X, tile.Y)
		tile.MoistureValue = moistureData.GetNormalized(tile.X, tile.Y)
		tile.CloudValue = cloudData.GetNormalized(tile.X, tile.Y)

		// initialize the tile
		tile.Init()
	})
}

func (textureGenerator *TextureGenerator) latLonToXYZ(lat, lon float64) (float64, float64, float64) {
	r := math.Cos(mathf.ToRadians(lon))
	x := r * math.Cos(mathf.ToRadians(lat))
	y := math.Sin(mathf.ToRadians(lon))
	z := r * math.Sin(mathf.ToRadians(lat))

	return x, y, z
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
		tile.Top = func() *textureset.Tile {
			return textureGenerator.getTopTile(tile)
		}
		tile.Bottom = func() *textureset.Tile {
			return textureGenerator.getBottomTile(tile)
		}
		tile.Left = func() *textureset.Tile {
			return textureGenerator.getLeftTile(tile)
		}
		tile.Right = func() *textureset.Tile {
			return textureGenerator.getRightTile(tile)
		}
	}

	return tile
}

func (textureGenerator *TextureGenerator) getTopTile(tile *textureset.Tile) *textureset.Tile {
	if tile.Y-1 > 0 {
		return textureGenerator.getTile(tile.X, tile.Y-1)
	}
	return nil
}

func (textureGenerator *TextureGenerator) getBottomTile(tile *textureset.Tile) *textureset.Tile {
	if tile.Y+1 < textureGenerator.height {
		return textureGenerator.getTile(tile.X, tile.Y+1)
	}
	return nil
}

func (textureGenerator *TextureGenerator) getLeftTile(tile *textureset.Tile) *textureset.Tile {
	return textureGenerator.getTile(tile.X-1, tile.Y)
}

func (textureGenerator *TextureGenerator) getRightTile(tile *textureset.Tile) *textureset.Tile {
	return textureGenerator.getTile(tile.X+1, tile.Y)
}

func (textureGenerator *TextureGenerator) forEachTile(fn func(tile *textureset.Tile)) {
	textureGenerator.forEachPixel(func(x, y int) {
		tile := textureGenerator.getTile(x, y)
		fn(tile)
	})
}

func (textureGenerator *TextureGenerator) imageFromEachTile(fn func(tile *textureset.Tile) *textureset.Color) *image.RGBA {
	rect := image.Rect(0, 0, textureGenerator.width, textureGenerator.height)
	img := image.NewRGBA(rect)

	textureGenerator.forEachTile(func(tile *textureset.Tile) {
		c := fn(tile)
		rgba := textureGenerator.color2RGBA(c)
		img.Set(tile.X, tile.Y, rgba)
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

func (textureGenerator *TextureGenerator) GenerateBumpMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.BumpColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateSpecularMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.SpecularColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateCloudMapTexture() *image.RGBA {
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return tile.CloudColor()
	})
	return img
}

func (textureGenerator *TextureGenerator) calculateNormalMapColor(tile *textureset.Tile, strength float64) *textureset.Color {
	left := textureGenerator.getTile(tile.X-1, tile.Y)
	if tile.X-1 < 0 {
		left = tile
	}

	right := tile
	if tile.X+1 < textureGenerator.width {
		right = textureGenerator.getTile(tile.X+1, tile.Y)
	}

	top := textureGenerator.getTile(tile.X, tile.Y-1)
	if tile.Y-1 < 0 {
		top = tile
	}

	bottom := tile
	if tile.Y+1 < textureGenerator.height {
		bottom = textureGenerator.getTile(tile.X, tile.Y+1)
	}

	leftColor := left.BumpColor()
	rightColor := right.BumpColor()
	topColor := top.BumpColor()
	bottomColor := bottom.BumpColor()

	xLeft := leftColor.Grayscale() * strength
	xRight := rightColor.Grayscale() * strength
	yTop := topColor.Grayscale() * strength
	yBottom := bottomColor.Grayscale() * strength

	xDelta := ((xLeft - xRight) + 1.0) * 0.5
	yDelta := ((yTop - yBottom) + 1.0) * 0.5

	return textureset.NewRGBAColor(xDelta, yDelta, 1.0, 1.0)
}

func (textureGenerator *TextureGenerator) GenerateNormalMapTexture(strength float64) *image.RGBA {
	strength = mathf.Clamp(strength, 0.0, 10.0)
	img := textureGenerator.imageFromEachTile(func(tile *textureset.Tile) *textureset.Color {
		return textureGenerator.calculateNormalMapColor(tile, strength)
	})
	return img
}

func (textureGenerator *TextureGenerator) GenerateTextures(strength float64) *image.RGBA {
	rect := image.Rect(0, 0, textureGenerator.width*2.0, textureGenerator.height*2.0)
	img := image.NewRGBA(rect)

	textureGenerator.forEachTile(func(tile *textureset.Tile) {
		biomeColor := tile.BiomeColor()
		normalColor := textureGenerator.calculateNormalMapColor(tile, strength)
		specularColor := tile.SpecularColor()
		cloudColor := tile.CloudColor()

		normalX := tile.X + textureGenerator.width
		normalY := tile.Y + textureGenerator.height

		specularX := tile.X + textureGenerator.width
		specularY := tile.Y

		cloudX := tile.X
		cloudY := tile.Y + textureGenerator.height

		img.Set(tile.X, tile.Y, textureGenerator.color2RGBA(biomeColor))
		img.Set(cloudX, cloudY, textureGenerator.color2RGBA(cloudColor))
		img.Set(specularX, specularY, textureGenerator.color2RGBA(specularColor))
		img.Set(normalX, normalY, textureGenerator.color2RGBA(normalColor))
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
