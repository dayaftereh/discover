package orbit_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/dayaftereh/discover/server/mathf"

	"github.com/dayaftereh/discover/server/mathf/orbit"
)

// Tool for simulation plots: https://plot.ly/create/

func TestOrbitSimulationPerfectCircle(t *testing.T) {
	centralBodyMass := 10.0   // kg
	centralBodyRadius := 10.0 // m
	mu := mathf.GravitationalConstant * centralBodyMass
	apogee := 30.0                // m
	perigee := 25.0               // m
	inclination := 92.5           // deg
	rightAscension := 157.6       // deg
	argumentOfPeriapsis := 210.42 // deg

	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		Apogee:              &apogee,
		Perigee:             &perigee,
		Inclination:         &inclination,
		RightAscension:      &rightAscension,
		ArgumentOfPeriapsis: &argumentOfPeriapsis,
	})

	Simulate(t, orbit, "perfect-circle")
}

func TestOrbitSimulationEllipticOrbit(t *testing.T) {
	centralBodyMass := 10.0   // kg
	centralBodyRadius := 10.0 // m
	mu := mathf.GravitationalConstant * centralBodyMass
	apogee := 69.0  // m
	perigee := 25.0 // m
	eccentricity := 0.7
	inclination := 92.5           // deg
	rightAscension := 157.6       // deg
	argumentOfPeriapsis := 210.42 // deg

	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		Apogee:              &apogee,
		Perigee:             &perigee,
		Eccentricity:        &eccentricity,
		Inclination:         &inclination,
		RightAscension:      &rightAscension,
		ArgumentOfPeriapsis: &argumentOfPeriapsis,
	})

	Simulate(t, orbit, "elliptic-orbit")
}

func TestOrbitSimulationEllipticOrbitPe(t *testing.T) {
	centralBodyMass := 10.0   // kg
	centralBodyRadius := 10.0 // m
	
	mu := mathf.GravitationalConstant * centralBodyMass
	apogee := 69.0  // m
	perigee := 25.0 // m
	eccentricity := 0.7
	inclination := 92.5           // deg
	rightAscension := 157.6       // deg
	argumentOfPeriapsis := 210.42 // deg

	orbit := orbit.OrbitFromParams(&orbit.OrbitParameter{
		MU:                  &mu,
		CentralBodyRadius:   &centralBodyRadius,
		Apogee:              &apogee,
		Perigee:             &perigee,
		Eccentricity:        &eccentricity,
		Inclination:         &inclination,
		RightAscension:      &rightAscension,
		ArgumentOfPeriapsis: &argumentOfPeriapsis,
	})

	Simulate(t, orbit, "elliptic-orbit")
}

func Simulate(t *testing.T, orbit *orbit.Orbit, name string) {
	// get period and position
	period := orbit.Period()
	position := orbit.Position()

	t.Logf("Period: %f s", period)
	t.Logf("Postion: %v", position)

	// create the output data
	data := "X;Y;Z\n"
	// write start position
	data = fmt.Sprintf("%s%f;%f;%f\n", data, position.X, position.Y, position.Z)

	accurent := 1000
	for i := 1; i < (accurent + 2); i++ {
		// get time step in seconds
		dt := float64(i) * (period / (float64(accurent)))
		// update to new orbit
		newOrbit := orbit.Update(dt)
		// get new position
		position := newOrbit.Position()
		t.Logf("(%d) Time: %f s - Position: %v", i, dt, position)
		// append the position to data
		data = fmt.Sprintf("%s%f;%f;%f\n", data, position.X, position.Y, position.Z)
	}

	fileName := fmt.Sprintf("./data-%s.csv", name)

	// data to bytes
	bytes := []byte(data)
	// write the file
	err := ioutil.WriteFile(fileName, bytes, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
