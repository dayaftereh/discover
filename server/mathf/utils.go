package mathf

import "math"

func ToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func ToDegress(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

func CloseZero(x float64) bool {
	return math.Abs(x) < Epsilon
}
