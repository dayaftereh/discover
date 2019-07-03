package data

type Game struct {
	Players     map[string]*Player
	StarSystems map[int64]*StarSystem
}
