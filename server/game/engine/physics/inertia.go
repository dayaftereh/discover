package physics

import (
	"github.com/dayaftereh/discover/server/mathf"
)

// https://en.wikipedia.org/wiki/List_of_moments_of_inertia

func CalculateSphereInertia(radius float64, mass float64) *mathf.Vec3 {
	I := 2.0 / 5.0 * mass * (radius * radius)
	inertia := mathf.NewVec3(I, I, I)
	return inertia
}

func CalculateCuboidInertia(width float64, height float64, depth float64, mass float64) *mathf.Vec3 {
	Ix := 1.0 / 12.0 * mass * (height*height + depth*depth)
	Iy := 1.0 / 12.0 * mass * (width*width + depth*depth)
	Iz := 1.0 / 12.0 * mass * (width*width + height*height)
	inertia := mathf.NewVec3(Ix, Iy, Iz)
	return inertia
}
