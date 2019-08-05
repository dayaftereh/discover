package common

type Status struct {
	ID            string  `json:"id"`
	Authenticated bool    `json:"authenticated"`
	Name          *string `json:"name"`
	StarSystem    *string `json:"starSystem"`
	Admin         bool    `json:"admin"`
}

func NewStatus() *Status {
	return &Status{}
}
