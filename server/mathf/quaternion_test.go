package mathf_test

import (
	"math"
	"testing"

	"github.com/dayaftereh/discover/server/mathf"
	"github.com/stretchr/testify/assert"
)

func AssertQuaternionCloseEqual(t *testing.T, q *mathf.Quaternion, x float64, y float64, z float64, w float64) {
	assert.InDelta(t, x, q.X, mathf.Epsilon)
	assert.InDelta(t, y, q.Y, mathf.Epsilon)
	assert.InDelta(t, z, q.Z, mathf.Epsilon)
	assert.InDelta(t, w, q.W, mathf.Epsilon)
}

func TestQuaternionNewQuaternion(t *testing.T) {
	q := mathf.NewQuaternion(42.0, 31.3, -7, 4)
	AssertQuaternionCloseEqual(t, q, 42.0, 31.3, -7, 4)
}

func TestQuaternionNewZeroQuaternion(t *testing.T) {
	q := mathf.NewZeroQuaternion()
	AssertQuaternionCloseEqual(t, q, 0, 0, 0, 1.0)
}

func TestQuaternionSet(t *testing.T) {
	q := mathf.NewZeroQuaternion()
	q.Set(42.0, 31.3, -7, 4)
	AssertQuaternionCloseEqual(t, q, 42.0, 31.3, -7, 4)
}

func TestQuaternionFromVectorsX(t *testing.T) {
	u := mathf.NewVec3(1, 0, 0)
	v := mathf.NewVec3(-1, 0, 0)

	q := mathf.QuaternionFromVectors(u, v)
	r := q.MultiplyVec(mathf.NewUnitX())

	AssertVec3CloseEqual(t, r, -1, 0, 0)
}

func TestQuaternionFromVectorsY(t *testing.T) {
	u := mathf.NewVec3(0, 1, 0)
	v := mathf.NewVec3(0, -1, 0)

	q := mathf.QuaternionFromVectors(u, v)
	r := q.MultiplyVec(mathf.NewUnitY())

	AssertVec3CloseEqual(t, r, 0, -1, 0)
}

func TestQuaternionToEulerX(t *testing.T) {
	q := mathf.QuaternionFromAxisAngle(mathf.NewUnitX(), math.Pi/4.0)
	e := q.ToEuler()

	AssertVec3CloseEqual(t, e, math.Pi/4.0, 0, 0)
}

func TestQuaternionToEulerY(t *testing.T) {
	q := mathf.QuaternionFromAxisAngle(mathf.NewUnitY(), math.Pi/4.0)
	e := q.ToEuler()

	AssertVec3CloseEqual(t, e, 0, math.Pi/4.0, 0)
}

func TestQuaternionToEulerZ(t *testing.T) {
	q := mathf.QuaternionFromAxisAngle(mathf.NewUnitZ(), math.Pi/4.0)
	e := q.ToEuler()

	AssertVec3CloseEqual(t, e, 0, 0, math.Pi/4.0)
}

func TestQuaternionFromVectorsZ(t *testing.T) {
	u := mathf.NewVec3(0, 0, 1)
	v := mathf.NewVec3(0, 0, -1)

	q := mathf.QuaternionFromVectors(u, v)
	r := q.MultiplyVec(mathf.NewUnitZ())

	AssertVec3CloseEqual(t, r, 0, 0, -1)
}

func TestQuaternionSlerp(t *testing.T) {
	qa := mathf.QuaternionFromAxisAngle(mathf.NewUnitZ(), math.Pi/4)
	qb := mathf.QuaternionFromAxisAngle(mathf.NewUnitZ(), -math.Pi/4)

	r := qa.Slerp(qb, 0.5)

	AssertQuaternionCloseEqual(t, r, 0, 0, 0, 1)
}

func TestQuaternionSlerpZero(t *testing.T) {
	qa := mathf.NewZeroQuaternion()
	qb := mathf.NewZeroQuaternion()

	r := qa.Slerp(qb, 0.5)

	AssertQuaternionCloseEqual(t, r, 0, 0, 0, 1)
}

func TestQuaternionInverse(t *testing.T) {
	q := mathf.NewQuaternion(1, 2, 3, 4)
	denominator := 1.0*1.0 + 2.0*2.0 + 3.0*3.0 + 4.0*4.0

	i := q.Inverse()

	assert.InDelta(t, -1.0/denominator, i.X, mathf.Epsilon)
	assert.InDelta(t, -2.0/denominator, i.Y, mathf.Epsilon)
	assert.InDelta(t, -3.0/denominator, i.Z, mathf.Epsilon)
	assert.InDelta(t, 4.0/denominator, i.W, mathf.Epsilon)
}
