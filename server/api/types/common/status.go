package common

type Status struct {
	Id   string  `json:"id"`
	Name *string `json:"name"`
}

func NewStatus() *Status {
	return &Status{}
}
