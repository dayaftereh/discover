package mathf_test

import (
	"math"
	"testing"

	"github.com/dayaftereh/discover/server/mathf"
	"github.com/stretchr/testify/assert"
)

func AssertVec3CloseEqual(t *testing.T, v *mathf.Vec3, x float64, y float64, z float64) {
	assert.InDelta(t, x, v.X, mathf.Epsilon)
	assert.InDelta(t, y, v.Y, mathf.Epsilon)
	assert.InDelta(t, z, v.Z, mathf.Epsilon)
}

func TestVec3NewVec3(t *testing.T) {
	v := mathf.NewVec3(42.0, 31.3, 7)
	AssertVec3CloseEqual(t, v, 42.0, 31.3, 7)
}

func TestVec3NewZeroVec3(t *testing.T) {
	v := mathf.NewZeroVec3()
	AssertVec3CloseEqual(t, v, 0, 0, 0)
}

func TestVec3NewUnitX(t *testing.T) {
	v := mathf.NewUnitX()
	AssertVec3CloseEqual(t, v, 1, 0, 0)
}

func TestVec3NewUnitY(t *testing.T) {
	v := mathf.NewUnitY()
	AssertVec3CloseEqual(t, v, 0, 1, 0)
}

func TestVec3NewUnitZ(t *testing.T) {
	v := mathf.NewUnitZ()
	AssertVec3CloseEqual(t, v, 0, 0, 1)
}

func TestVec3Set(t *testing.T) {
	v := mathf.NewZeroVec3()
	v.Set(42.0, 31.3, 7)
	AssertVec3CloseEqual(t, v, 42.0, 31.3, 7)
}

func TestVec3SetVec(t *testing.T) {
	v1 := mathf.NewZeroVec3()
	v2 := mathf.NewVec3(42.0, 31.3, 7)
	v := v1.SetVec(v2)
	AssertVec3CloseEqual(t, v, 42.0, 31.3, 7)
}

func TestVec3Add(t *testing.T) {
	v1 := mathf.NewVec3(82, 97.5, -23.77)
	v2 := mathf.NewVec3(42.0, -31.3, 7)

	v := v1.AddVec(v2)

	AssertVec3CloseEqual(t, v, (42.0 + 82), (97.5 + (-31.3)), ((-23.77) + 7))
}

func TestVec3Subtract(t *testing.T) {
	v1 := mathf.NewVec3(82, 97.5, -23.77)
	v2 := mathf.NewVec3(42.0, -31.3, 7)

	v := v1.SubtractVec(v2)

	AssertVec3CloseEqual(t, v, (82 - 42), (97.5 - (-31.3)), ((-23.77) - 7))
}

func TestVec3Normalize(t *testing.T) {
	v := mathf.NewVec3(82, 97.5, -23.77)
	v.Normalize()
	l := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	assert.InDelta(t, 1, l, mathf.Epsilon)
}

func TestVec3NormalizeZero(t *testing.T) {
	v := mathf.NewZeroVec3()
	v.Normalize()
	l := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	assert.InDelta(t, 0, l, mathf.Epsilon)
}

func TestVec3Unit(t *testing.T) {
	v := mathf.NewVec3(82, 97.5, -23.77)
	n := v.Unit()
	l := math.Sqrt(n.X*n.X + n.Y*n.Y + n.Z*n.Z)
	assert.InDelta(t, 1, l, mathf.Epsilon)
}

func TestVec3UnitZero(t *testing.T) {
	v := mathf.NewZeroVec3()
	n := v.Unit()
	l := math.Sqrt(n.X*n.X + n.Y*n.Y + n.Z*n.Z)
	assert.InDelta(t, 0, l, mathf.Epsilon)
}

func TestVec3Length(t *testing.T) {
	v := mathf.NewVec3(42, 31.3, -7)
	l1 := v.Length()
	l2 := math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	assert.InDelta(t, l2, l1, mathf.Epsilon)
}

func TestVec3SqrtLength(t *testing.T) {
	v := mathf.NewVec3(42, 31.3, -7)
	l1 := v.SqrtLength()
	l2 := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	assert.InDelta(t, l2, l1, mathf.Epsilon)
}

func TestVec3DistanceTo(t *testing.T) {
	v1 := mathf.NewVec3(42, 31.3, -7)
	v2 := mathf.NewVec3(92, -25.3, 19.1)
	d1 := v1.DistanceTo(v2)
	d2 := 79.90475580339383
	assert.InDelta(t, d2, d1, mathf.Epsilon)
}

func TestVec3SqrtDistanceTo(t *testing.T) {
	v1 := mathf.NewVec3(42, 31.3, -7)
	v2 := mathf.NewVec3(92, -25.3, 19.1)
	d1 := v1.SqrtDistanceTo(v2)
	d2 := 79.90475580339383
	assert.InDelta(t, d2*d2, d1, mathf.Epsilon)
}

func TestVec3Multiply(t *testing.T) {
	v := mathf.NewVec3(42, 31.3, -7)
	r := v.Multiply(15.4)
	AssertVec3CloseEqual(t, r, (v.X * 15.4), (v.Y * 15.4), (v.Z * 15.4))
}

func TestVec3MultiplyVec(t *testing.T) {
	v1 := mathf.NewVec3(42, 31.3, -7)
	v2 := mathf.NewVec3(32.5, -19.3, 8)
	r := v1.MultiplyVec(v2)
	AssertVec3CloseEqual(t, r, (v1.X * v2.X), (v1.Y * v2.Y), (v1.Z * v2.Z))
}

func TestVec3Dot(t *testing.T) {
	v1 := mathf.NewVec3(1, 2, 3)
	v2 := mathf.NewVec3(4, 5, 6)
	dot := v1.Dot(v2)
	assert.InDelta(t, 4+10+18, dot, mathf.Epsilon)
}

func TestVec3Cross(t *testing.T) {
	v1 := mathf.NewVec3(1, 2, 3)
	v2 := mathf.NewVec3(4, 5, 6)
	c := v1.Cross(v2)
	AssertVec3CloseEqual(t, c, -3, 6, -3)
}

func TestVec3Negate(t *testing.T) {
	v := mathf.NewVec3(1, -2, 3)
	n := v.Negate()
	AssertVec3CloseEqual(t, n, -1, 2, -3)
}

func TestVec3Tangents(t *testing.T) {

}

func TestVec3Clone(t *testing.T) {
	v := mathf.NewVec3(1, -2, 3)
	c := v.Clone()
	AssertVec3CloseEqual(t, c, 1, -2, 3)
}

func TestVec3Lerp(t *testing.T) {
}

func TestVec3AlmostEquals(t *testing.T) {
	v1 := mathf.NewVec3(1, -2.77, 31.2)
	v2 := mathf.NewVec3(1, -2.77, 31.2)
	e := v1.AlmostEquals(v2, mathf.Epsilon)
	assert.True(t, e)
}

func TestVec3AlmostEqualsFalse(t *testing.T) {
	v1 := mathf.NewVec3(1, -2.77, 31.2)
	v2 := mathf.NewVec3(1, 2.77, 31.2)
	e := v1.AlmostEquals(v2, mathf.Epsilon)
	assert.False(t, e)
}

func TestVec3AlmostZero(t *testing.T) {
	v := mathf.NewZeroVec3()
	e := v.AlmostZero(mathf.Epsilon)
	assert.True(t, e)
}

func TestVec3AlmostZeroFalse(t *testing.T) {
	v := mathf.NewVec3(1, 2, 3)
	e := v.AlmostZero(mathf.Epsilon)
	assert.False(t, e)
}

func TestVec3IsAntiparallelTo(t *testing.T) {
	v1 := mathf.NewVec3(1, 0, 0)
	v2 := mathf.NewVec3(-1, 0, 0)
	e := v1.IsAntiparallelTo(v2, mathf.Epsilon)
	assert.True(t, e)
}

func TestVec3IsAntiparallelToFalse(t *testing.T) {
	v1 := mathf.NewVec3(1, 0, -1)
	v2 := mathf.NewVec3(-1, 0, 0)
	e := v1.IsAntiparallelTo(v2, mathf.Epsilon)
	assert.False(t, e)
}
