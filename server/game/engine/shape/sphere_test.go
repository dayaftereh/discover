package shape_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/game/engine/shape"
)

func TestShpere(t *testing.T) {
	sphere := shape.NewSphere(123, 42)

	shapes := []shape.Shape{
		sphere,
	}

	t.Fatalf("lll:%d -- %f\n", sphere.ID(), sphere.Volume())
	t.Fatalf("lll:%d\n", len(shapes))
}
