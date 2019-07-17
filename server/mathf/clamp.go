package mathf

import "math"

func Clamp(v float64, min float64, max float64) float64 {
	return math.Min(math.Max(v, min), max)
}

func ClampVec3(vec *Vec3, min float64, max float64) *Vec3 {
	return NewVec3(
		Clamp(vec.X, min, max),
		Clamp(vec.Y, min, max),
		Clamp(vec.Z, min, max),
	)
}
