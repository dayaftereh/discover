package mathf_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/mathf"
)

func PrintOrbit(t *testing.T, i int, orbit *mathf.Orbit) {
	position := orbit.Position()
	velocity := orbit.Velocity()

	t.Logf("step: %d", i)
	t.Logf("position: %v", position)
	t.Logf("velocity: %v", velocity)
}

func TestOrbit(t *testing.T) {
	radius := mathf.NewVec3(0, 3000000, 0)
	velocity := mathf.NewVec3(-7000, 0, 0)
	orbit := mathf.NewOrbitFromVectors(900000000, radius, velocity)

	t.Logf("Orbit: %v\n", orbit)

	for i := 1; i < 10; i++ {
		t.Log("Before")
		orbit.Update(10.0)
		t.Log("After")
		t.Logf("Orbit: %v\n", orbit)
	}

}
