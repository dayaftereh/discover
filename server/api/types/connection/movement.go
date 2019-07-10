package connection

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Movement struct {
	Move     *mathf.Vec3 `json:"move"`
	Rotation *mathf.Vec3 `json:"rotation"`
}
