package data

type Universe struct {
	InitialStarSystem *int64
	StarSystems       map[int64]*StarSystem
}
