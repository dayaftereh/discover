package noise_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/texture/noise"
)

func TestImplicitGradientGetD2(t *testing.T) {

	implicitGradient := noise.NewImplicitGradient(1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1)

	WriteImage(t, implicitGradient, "implicit-gradient-d2")
}
