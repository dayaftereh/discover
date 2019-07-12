package mathf_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/mathf"
	"github.com/stretchr/testify/assert"
)

func AssertMat3CloseEqual(t *testing.T, m *mathf.Mat3,
	e1 float64, e2 float64, e3 float64,
	e4 float64, e5 float64, e6 float64,
	e7 float64, e8 float64, e9 float64) {

	assert.InDelta(t, e1, m.Element(0), mathf.Epsilon)
	assert.InDelta(t, e2, m.Element(1), mathf.Epsilon)
	assert.InDelta(t, e3, m.Element(2), mathf.Epsilon)

	assert.InDelta(t, e4, m.Element(3), mathf.Epsilon)
	assert.InDelta(t, e5, m.Element(4), mathf.Epsilon)
	assert.InDelta(t, e6, m.Element(5), mathf.Epsilon)

	assert.InDelta(t, e7, m.Element(6), mathf.Epsilon)
	assert.InDelta(t, e8, m.Element(7), mathf.Epsilon)
	assert.InDelta(t, e9, m.Element(8), mathf.Epsilon)

}

func TestMat3NewZeroMat3(t *testing.T) {
	m := mathf.NewZeroMat3()
	AssertMat3CloseEqual(t, m, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
}

func TestMat3NewIdentityMat3(t *testing.T) {
	m := mathf.NewIdentityMat3()
	AssertMat3CloseEqual(t, m, 1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 1.0)
}

func TestMat3MultiplyVec(t *testing.T) {
	m := mathf.NewMat3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)
	v := mathf.NewVec3(2, 3, 7)
	r := m.MultiplyVec(v)
	AssertVec3CloseEqual(t, r, 29.0, 65.0, 101.0)
}

func TestMat3Multiply(t *testing.T) {
	m1 := mathf.NewMat3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)
	m2 := mathf.NewMat3(5.0, 2.0, 4.0, 4.0, 5.0, 1.0, 1.0, 8.0, 0.0)
	r := m1.Multiply(m2)
	AssertMat3CloseEqual(t, r, 16, 36, 6, 46, 81, 21, 76, 126, 36)
}

func TestMat3Transpose(t *testing.T) {
	m := mathf.NewMat3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)
	r := m.Transpose()
	AssertMat3CloseEqual(t, r, 1, 4, 7, 2, 5, 8, 3, 6, 9)
}

func TestMat3Scale(t *testing.T) {
	m := mathf.NewMat3(1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0)
	v := mathf.NewVec3(1, 2, 3)
	r := m.Scale(v)
	AssertMat3CloseEqual(t, r, 1, 2, 3, 1, 2, 3, 1, 2, 3)
}

func TestMat3Mat3FromQuaternionZero(t *testing.T) {
	q := mathf.NewZeroQuaternion()
	m := mathf.Mat3FromQuaternion(q)
	v := mathf.NewVec3(1, 2, 3)
	r := m.MultiplyVec(v)
	AssertVec3CloseEqual(t, r, 1, 2, 3)
}

func TestMat3NewFromQuaternion(t *testing.T) {
	euler := mathf.NewVec3(0.222, 0.123, 1.234)
	q := mathf.QuaternionFromEuler(euler)
	m := mathf.Mat3FromQuaternion(q)

	v := mathf.NewVec3(1, 2, 3)
	mr := m.MultiplyVec(v)
	qr := q.MultiplyVec(v)

	AssertVec3CloseEqual(t, mr, qr.X, qr.Y, qr.Z)
}
