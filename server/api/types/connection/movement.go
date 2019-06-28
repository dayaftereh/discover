package connection

type Movement struct {
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Z            float64 `json:"z"`
	Acceleration int     `json:"acceleration"`
}
