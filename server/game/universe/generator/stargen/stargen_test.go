package stargen_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/universe/generator/stargen"
)

func TestStargen(t *testing.T) {
	sun, _ := stargen.GenerateStellarSystem(true, true, true)
	t.Log(sun, "f")
	/*for _, planet := range planets {
		t.Log(planet)
	}*/

}
