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

func WriteImage(t *testing.T, img *image.RGBA, name string) {
	filename := fmt.Sprintf("./textures-%s.png", name)
	fd, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(fd, img)
	if err != nil {
		t.Fatal(err)
	}
}

func ExportTextures(t *testing.T, textureGenerator *texture.TextureGenerator, name string) {

	textureGenerator.Init()

	textures := textureGenerator.GenerateTextures(3.0)

	filename := fmt.Sprintf("./textures-%s.png", name)

	fd, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(fd, textures)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTerrestrialTextureGenerator(t *testing.T) {
	width, height := 1024, 1024
	seed := utils.RandInt64(0, 1e10)
	textureSet := terrestrial.NewTerrestrialTextureSet()
	textureGenerator := texture.NewTextureGenerator(width, height, textureSet, seed)

	ExportTextures(t, textureGenerator, "terrestrial")
}
