package stargen_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
)

func TestStargen(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, planets := stargen.GenerateStellarSystem(true, true, true)
		//t.Log(sun)
		for _, planet := range planets {
			t.Log(planet)
		}
	}
}
