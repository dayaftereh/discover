package shape

import "github.com/dayaftereh/discover/server/mathf"

type Shape interface {
	ID() int64
	Volume() float64
	CalculateLocalInertia(mass float64) *mathf.Vec3
	BoundingSphereRadius() float64
	UpdateBoundingSphereRadius()
}

type internalShape struct {
	id                   int64
	boundingSphereRadius float64
}

func (shape *internalShape) ID() int64 {
	return shape.id
}

func (shape *internalShape) BoundingSphereRadius() float64 {
	return shape.boundingSphereRadius
}
