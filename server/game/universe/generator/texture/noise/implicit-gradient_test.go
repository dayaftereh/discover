package noise_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
)

func TestImplicitGradientGetD4(t *testing.T) {

	implicitGradient := noise.NewImplicitGradient(1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1)

	WriteImage4D(t, implicitGradient, "implicit-gradient-4d")
}

func TestImplicitGradientGetD2(t *testing.T) {

	implicitGradient := noise.NewImplicitGradient(1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1)

	WriteImage2D(t, implicitGradient, "implicit-gradient-2d")
}
