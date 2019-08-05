package types

type Player struct {
	Name       string  `json:"name"`
	StarSystem *string `json:"starSystem"`
	Admin      bool    `json:"admin"`
}
