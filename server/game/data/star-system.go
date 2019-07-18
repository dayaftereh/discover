package data

type StarSystem struct {
	ID      int64
	Name    *string
	Sun     *Sun
	Planets []Planet
}
