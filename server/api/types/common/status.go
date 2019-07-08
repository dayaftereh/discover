package common

type Status struct {
	ID            string  `json:"id"`
	Authenticated bool    `json:"authenticated"`
	Name          *string `json:"name"`
	StarSystem    *int64  `json:"starSystem"`
}

func NewStatus() *Status {
	return &Status{}
}
