package types

type Universe struct {
	InitialStarSystem *string             `json:"initialStarSystem"`
	StarSystems       map[string][]string `json:"starSystems"`
}
