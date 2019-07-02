package shape

import (
	"math"

	"github.com/dayaftereh/discover/server/mathf"
)

type Sphere struct {
	internalShape
	radius float64
}

func NewSphere(id int64, radius float64) *Sphere {
	return &Sphere{
		internalShape: internalShape{
			id: id,
		},
		radius: radius,
	}
}

func (sphere *Sphere) Volume() float64 {
	return 4.0 * math.Pi * sphere.radius / 3.0
}

func (sphere *Sphere) CalculateLocalInertia(mass float64) *mathf.Vec3 {
	i := 2.0 * mass * sphere.radius * sphere.radius / 5.0
	return mathf.NewVec3(i, i, i)
}

func (sphere *Sphere) UpdateBoundingSphereRadius() {
	sphere.boundingSphereRadius = sphere.radius
}
