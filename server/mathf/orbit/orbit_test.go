package orbit_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/mathf"
	"github.com/stretchr/testify/assert"

	"github.com/dayaftereh/discover/server/mathf/orbit"
)

func AssertVec3CloseEqual(t *testing.T, v *mathf.Vec3, x float64, y float64, z float64) {
	assert.InDelta(t, x, v.X, mathf.Epsilon)
	assert.InDelta(t, y, v.Y, mathf.Epsilon)
	assert.InDelta(t, z, v.Z, mathf.Epsilon)
}

func DefaultOrbit() *orbit.Orbit {
	r := mathf.NewVec3(-6045, -3490, 2500)
	v := mathf.NewVec3(-3.457, 6.618, 2.533)
	mu := 398600.0              // earth mu
	centralBodyRadius := 6371.0 // earth radius
	return orbit.NewOrbit(r, v, mu, centralBodyRadius)
}

func TestOrbitAngularMomentum(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.AngularMomentum()
	AssertVec3CloseEqual(t, actual, -25385.17, 6669.485, -52070.74)
}

func TestOrbitRadialVelocity(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.RadialVelocity()
	assert.InDelta(t, 0.55746792, actual, mathf.Epsilon)
}

func TestOrbitEccentricity(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.Eccentricity()
	AssertVec3CloseEqual(t, actual, -0.091604856046, -0.1422073715, 0.0264439282)
}

func TestOrbitSemimajorAxis(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.SemimajorAxis()
	assert.InDelta(t, 8788.0951173776, actual, mathf.Epsilon)
}

func TestOrbitSemiminorAxis(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.SemiminorAxis()
	assert.InDelta(t, 8658.331432693347, actual, mathf.Epsilon)
}

func TestOrbitSemilatusRectum(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.SemilatusRectum()
	assert.InDelta(t, 8530.483818970712, actual, mathf.Epsilon)
}

func TestOrbitInclination(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.Inclination()
	assert.InDelta(t, 153.24922851824746, actual, mathf.Epsilon)
}

func TestOrbitNodeLine(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.NodeLine()
	AssertVec3CloseEqual(t, actual, -6669.485, -25385.17, 0.0)
}

func TestOrbitRightAscension(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.RightAscension()
	assert.InDelta(t, 255.27928533439618, actual, mathf.Epsilon)
}

func TestOrbitArgumentOfPeriapsis(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.ArgumentOfPeriapsis()
	assert.InDelta(t, 20.06832, actual, mathf.Epsilon)
}

func TestOrbitTrueAnomaly(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.TrueAnomaly()
	assert.InDelta(t, 28.445628306614964, actual, mathf.Epsilon)
}

func TestOrbitApoapsis(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.Apoapsis()
	assert.InDelta(t, 10292.725501794834, actual, mathf.Epsilon)
}

func TestOrbitPeriapsis(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.Periapsis()
	assert.InDelta(t, 7283.464732960478, actual, mathf.Epsilon)
}

func TestOrbitPeriod(t *testing.T) {
	orbit := DefaultOrbit()

	actual := orbit.Period()
	assert.InDelta(t, 8198.857616829207, actual, mathf.Epsilon)
}

func UpdateableOrbit() *orbit.Orbit {
	r := mathf.NewVec3(7000, -12124, 0)
	v := mathf.NewVec3(2.6679, 4.6210, 0)
	mu := 398600.0              // earth mu
	centralBodyRadius := 6371.0 // earth radius

	return orbit.NewOrbit(r, v, mu, centralBodyRadius)
}

func TestOrbitUniversalAnomaly(t *testing.T) {
	orbit := UpdateableOrbit()

	actual := orbit.UniversalAnomaly(3600)
	assert.InDelta(t, 253.53449076412875, actual, mathf.Epsilon)
}

func TestOrbitUpdate(t *testing.T) {
	orbit := UpdateableOrbit()

	actual := orbit.Update(3600)

	actualR := actual.Position()
	actualV := actual.Velocity()

	AssertVec3CloseEqual(t, actualR, -3297.7686251992877, 7413.396645787411, 0.0)
	AssertVec3CloseEqual(t, actualV, -8.297603024266518, -0.964044944673768, -0.0)

}
