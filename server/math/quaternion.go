package math

import "math"

type Quaternion struct {
	X float64
	Y float64
	Z float64
	W float64
}

// FromAxisAngle set the quaternion components to the given axis and angle. angle in radians.
func (quaternion *Quaternion) FromAxisAngle(axis Vec3, angle float64) *Quaternion {
	s := math.Sin(angle * 0.5)

	quaternion.X = axis.X * s
	quaternion.Y = axis.Y * s
	quaternion.Z = axis.Z * s
	quaternion.W = math.Cos(angle * 0.5)

	return quaternion
}

func (quaternion *Quaternion) FromVec(u Vec3, v Vec3) *Quaternion {
	if u.IsAntiparallelTo(v) {
		t1 := Vec3{}
		t2 := Vec3{}

		u.Tangents(t1, t2)
		quaternion.FromAxisAngle(t1, math.Pi)
		return quaternion
	}

	a := u.Cross(v)

	quaternion.X = a.X
	quaternion.Y = a.Y
	quaternion.Z = a.Z
	quaternion.W = math.Sqrt(u.SqrtLength()*v.SqrtLength()) + u.Dot(v)

	quaternion.Normalize()

	return quaternion
}

func (quaternion *Quaternion) Normalize() *Quaternion {

}

func (quaternion *Quaternion) Multiply(other Quaternion) *Quaternion {

}

func (quaternion *Quaternion) Inverse() *Quaternion {

}

func (quaternion *Quaternion) Conjugate() *Quaternion {

}

func (quaternion *Quaternion) MultiplyVec(vec Vec3) *Quaternion {

}

func (quaternion *Quaternion) Copy() *Quaternion {

}

func (quaternion *Quaternion) ToEuler() Vec3 {

}

func (quaternion *Quaternion) FromEuler(vec Vec3) *Quaternion {

}

func (quaternion *Quaternion) Integrate(angularVelocity Vec3, dt float64, angularFactor Vec3) *Quaternion {

}

func (quaternion *Quaternion) Slerp(toQuaternion Quaternion, t float64) *Quaternion {

}
