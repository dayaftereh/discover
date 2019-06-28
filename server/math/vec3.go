package math

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (vec *Vec3) Add(other Vec3) *Vec3 {
	return nil
}

func (vec *Vec3) Subtract(other Vec3) *Vec3 {
	return nil
}

func (vec *Vec3) Normalize() *Vec3 {
	return nil
}

func (vec *Vec3) Unit() *Vec3 {
	return nil
}

func (vec *Vec3) Length() float64 {

}

func (vec *Vec3) SqrtLength() float64 {

}

func (vec *Vec3) DistanceTo(other Vec3) float64 {

}

func (vec *Vec3) SqrtDistanceTo(other Vec3) float64 {

}

func (vec *Vec3) Multiply(scale float64) float64 {

}

func (vec *Vec3) MultiplyVec(other Vec3) float64 {

}

func (vec *Vec3) Dot(other Vec3) float64 {

}

func (vec *Vec3) Cross(other Vec3) *Vec3 {
	return nil
}

func (vec *Vec3) Negate() Vec3 {

}

func (vec *Vec3) Tangents(t1 Vec3, t2 Vec3) {}

func (vec *Vec3) Copy() Vec3 {}

func (vec *Vec3) Lerp(other Vec3, t float64) Vec3 {}

func (vec *Vec3) AlmostEquals(other Vec3, precision float64) bool {}

func (vec *Vec3) AlmostZero(other Vec3, precision float64) bool {}

func (vec *Vec3) IsAntiparallelTo(other Vec3) bool {
	return false
}
