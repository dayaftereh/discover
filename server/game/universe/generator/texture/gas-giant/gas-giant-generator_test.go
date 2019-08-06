package gasgiant_test

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"testing"

	gasgiant "github.com/dayaftereh/discover/server/game/universe/generator/texture/gas-giant"
	"github.com/dayaftereh/discover/server/utils"
)

func WriteImage(t *testing.T, img *image.RGBA, name string) {
	filename := fmt.Sprintf("./gas-giant-%s.png", name)
	fd, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	err = png.Encode(fd, img)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGasGiantGenerator(t *testing.T) {
	seed := utils.RandInt64(1, 1e6)

	linearGradient := RandomLinearGradient1D()

	gasGiantGenerator := gasgiant.NewGasGiantGenerator(linearGradient, 1024, 1024, seed)

	gasGiantGenerator.Init()

	img := gasGiantGenerator.GetFractalImage()

	WriteImage(t, img, "test")
}
