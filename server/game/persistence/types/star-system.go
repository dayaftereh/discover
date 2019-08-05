package types

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type StarSystem struct {
	Name          string      `json:"name"`
	Sun           *Sun        `json:"sun"`
	SpawnLocation *mathf.Vec3 `json:"spawnLocation"`
	Planets       []*Planet   `json:"planets"`
}
