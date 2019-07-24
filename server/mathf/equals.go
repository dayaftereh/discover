package mathf

import "math"

func CloseZero(x float64) bool {
	return math.Abs(x) < Epsilon
}

func CloseEquals(a float64, b float64) bool {
	return CloseZero(a - b)
}
