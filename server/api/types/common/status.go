package common

type Status struct {
	Id         string  `json:"id"`
	Name       *string `json:"name"`
	StarSystem *int64  `json:"starSystem"`
}

func NewStatus() *Status {
	return &Status{}
}
