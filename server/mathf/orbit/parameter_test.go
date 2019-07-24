package orbit_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/mathf/orbit"
)

func TestOrbitFromOrbitalElements(t *testing.T) {
	mu := 398600.0
	centralBodyRadius := 6371.0

	// values taken from Orbital Mechanics for Engineering Students, Example 4.7
	semilatusRectum := 16056.196688409433
	eccentricity := 1.4
	inclination := 30.0
	argumentOfPeriapsis := 60.0
	rightAscension := 40.0
	trueAnomaly := 30.0

	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		SemilatusRectum:     &semilatusRectum,
		Eccentricity:        &eccentricity,
		Inclination:         &inclination,
		ArgumentOfPeriapsis: &argumentOfPeriapsis,
		RightAscension:      &rightAscension,
		TrueAnomaly:         &trueAnomaly,
	})

	actualR := orbit.Position()
	actualV := orbit.Velocity()

	AssertVec3CloseEqual(t, actualR, -4039.895923201738, 4814.560480182377, 3628.6247021718837)
	AssertVec3CloseEqual(t, actualV, -10.385987618194685, -4.771921637340853, 1.743875)

}

func TestOrbitFromOrbitalElements2(t *testing.T) {
	mu := 398600.0
	centralBodyRadius := 6371.0

	// values taken from Orbital Mechanics for Engineering Students, Example 4.7
	apogee := 416.0
	perigee := 405.0
	inclination := 51.65
	rightAscension := 304.0847
	argumentOfPeriapsis := 117.7713

	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		Apogee:              &apogee,
		Perigee:             &perigee,
		Inclination:         &inclination,
		ArgumentOfPeriapsis: &argumentOfPeriapsis,
		RightAscension:      &rightAscension,
	})

	actualR := orbit.Position()
	actualV := orbit.Velocity()

	AssertVec3CloseEqual(t, actualR, 1311.5636463724552, 4699.598648496939, 4701.881415410206)
	AssertVec3CloseEqual(t, actualV, -5.641883624845521, 4.379639136504687, -2.803740788571159)

}
