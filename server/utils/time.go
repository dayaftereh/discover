package utils

import "time"

func SystemMillis() float64 {
	now := time.Now()
	unixNano := now.UnixNano()
	return float64(unixNano) / float64(time.Millisecond)
}

func SystemSeconds() float64 {
	now := time.Now()
	unixNano := now.UnixNano()
	return float64(unixNano) / float64(time.Second)
}
