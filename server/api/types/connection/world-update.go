package connection

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type GameObject struct {
	Type       *string     `json:"type"`
	Radius     *float64    `json:"radius"`
	Position   *mathf.Vec3 `json:"position"`
	Rotation   *mathf.Vec3 `json:"rotation"`
	Removeable *bool       `json:"removeable"`
	Color      *int64      `json:"color"`
}

type PlayerUpdate struct {
	GameObjectId *int64 `json:"gameObjectId"`
}

type WorldUpdateMessage struct {
	Type    MessageType           `json:"type"`
	Tick    *int64                `json:"tick"`
	Time    *float64              `json:"time"`
	Player  *PlayerUpdate         `json:"player"`
	Objects map[int64]*GameObject `json:"objects"`
}
