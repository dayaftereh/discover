package texture_test

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture"
	"github.com/dayaftereh/discover/server/game/universe/generator/texture/terrestrial"
)

func ExportTextures(t *testing.T, textureGenerator *texture.TextureGenerator, name string) {

	textureGenerator.Init()

	heightMap := textureGenerator.GenerateHeightMapTexture()
	heatMap := textureGenerator.GenerateHeatMapTexture()
	moistureMap := textureGenerator.GenerateMoistureMapTexture()
	biomeMap := textureGenerator.GenerateBiomeMapTexture()

	size := heightMap.Rect.Size()
	width := size.X * 2
	height := size.Y * 2

	output := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			heightColor := heightMap.At(x, y)
			heatColor := heatMap.At(x, y)
			moistureColor := moistureMap.At(x, y)
			biomeColor := biomeMap.At(x, y)

			heatX := x + size.X
			heatY := y

			moistureX := x
			moistureY := y + size.Y

			biomeX := x + size.X
			biomeY := y + size.Y

			output.Set(x, y, heightColor)
			output.Set(heatX, heatY, heatColor)
			output.Set(moistureX, moistureY, moistureColor)
			output.Set(biomeX, biomeY, biomeColor)
		}
	}

	filename := fmt.Sprintf("./textures-%s.png", name)

	fd, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(fd, output)
	if err != nil {
		//t.Fatal(err)
	}
}

func TestTerrestrialTextureGenerator(t *testing.T) {
	width, height := 256, 256
	seed := utils.RandInt64(0, 1000)
	textureSet := terrestrial.NewTerrestrialTextureSet()
	textureGenerator := texture.NewTextureGenerator(width, height, textureSet, seed)

	ExportTextures(t, textureGenerator, "terrestrial")
}
