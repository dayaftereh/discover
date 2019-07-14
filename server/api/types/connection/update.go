package connection

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type GameObject struct {
	Radius     *float64    `json:"radius"`
	Position   *mathf.Vec3 `json:"position"`
	Rotation   *mathf.Vec3 `json:"rotation"`
	Removeable *bool       `json:"removeable"`
}

type WorldUpdate struct {
	Type    MessageType           `json:"type"`
	Tick    *int64                `json:"tick"`
	Time    *float64              `json:"time"`
	Player  *GameObject           `json:"player"`
	Objects map[int64]*GameObject `json:"objects"`
}