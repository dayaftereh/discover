package object_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/engine/shape"
)

func TestShpere(t *testing.T) {
	shapes := []*shape.Shape{
		&shape.Shape{},
		shape.NewSphere(1, 1),
	}
}
