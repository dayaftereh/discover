package object

import (
	"log"

	"github.com/dayaftereh/discover/server/game/data"
	"github.com/dayaftereh/discover/server/mathf"
)

// GravitationalConstant is an empirical physical constant
var GravitationalConstant float64 = 1

type Planet struct {
	id              int64
	color           int64
	radius          float64
	sunMass         float64
	rigidbody       *RigidBody
	initialPosition *mathf.Vec3
	initialForce    *mathf.Vec3
	orbit           *mathf.Orbit2
}

func NewPlanet(id int64) *Planet {
	rigidbody := NewRigidBody(1)

	return &Planet{
		id:        id,
		rigidbody: rigidbody,
	}
}

func (planet *Planet) Load(sunMass float64, data *data.Planet) {
	planet.sunMass = sunMass

	// set planet
	planet.color = data.Color
	planet.radius = data.Radius
	planet.initialForce = data.Force
	planet.initialPosition = data.Position.Clone()

	// set rigidbody
	planet.rigidbody.Mass = data.Mass
	planet.rigidbody.Position = data.Position.Clone()

	planet.rigidbody.Velocity = mathf.NewVec3(113, -25, 0)

	// load the Inertia
	planet.rigidbody.Inertia = CalculateSphereInertia(planet.radius, planet.rigidbody.Mass)
	planet.rigidbody.UpdateInertiaWorld(true)
}

func (planet *Planet) ID() int64 {
	return planet.id
}

func (planet *Planet) Radius() float64 {
	return planet.radius
}

func (planet *Planet) RigidBody() *RigidBody {
	return planet.rigidbody
}

func (planet *Planet) force(distance *mathf.Vec3) *mathf.Vec3 {
	radius := distance.Length()

	force := distance.Multiply(-20000 * planet.rigidbody.Mass * 1.85).Divide(radius * radius * radius)

	return force
}

func (planet *Planet) Update(delta float64) {

	if delta > 0.2 {
		delta = 0
	}

	distance := planet.rigidbody.Position.Clone()

	p1 := planet.rigidbody.Position.Clone()
	v1 := planet.rigidbody.Velocity.Clone()
	force1 := planet.force(distance)

	log.Printf("force1: %v", force1)

	p2 := p1.AddVec(v1.Multiply(delta / 2.0))
	v2 := v1.AddVec(force1.Multiply(delta / 2.0))

	force2 := planet.force(p2)

	p3 := p1.AddVec(v2.Multiply(delta / 2.0))
	v3 := v1.AddVec(force2.Multiply(delta / 2.0))

	force3 := planet.force(p3)

	p4 := p1.AddVec(v3.Multiply(delta))
	v4 := v1.AddVec(force3.Multiply(delta))

	force4 := planet.force(p4)

	/*
		var force = this.force(distance);
		var fx = force.x;//-1000 * 1.85 * distance.x / Math.pow(distanceAbs,3);
		var fy = force.y;//-1000 * 1.85 * distance.y / Math.pow(distanceAbs,3);

		p1x = this.x;
		v1x = this.vx;
		a1x = fx;

		p1y = this.y;
		v1y = this.vy;
		a1y = fy;

		p2x = p1x + v1x * (dt/2);
		p2y = p1y + v1y * (dt/2);


		v2x = v1x + a1x * (dt/2);
		var force = this.force({x:p2x, y:p2y});
		var fx = force.x;//-1000 * 1.85 * distance.x / Math.pow(distanceAbs,3);
		var fy = force.y;//-1000 * 1.85 * distance.y / Math.pow(distanceAbs,3);
		a2x = fx;

		v2y = v1y + a1y * (dt/2);
		a2y = fy;

		p3x = p1x + v2x * (dt/2);
		p3y = p1y + v2y * (dt/2);
		var force = this.force({x:p3x, y:p3y});
		var fx = force.x;//-1000 * 1.85 * distance.x / Math.pow(distanceAbs,3);
		var fy = force.y;//-1000 * 1.85 * distance.y / Math.pow(distanceAbs,3);
		a3x = fx;
		a3y = fy;

		v3x = v1x + a2x * (dt/2);
		v3y = v1y + a2y * (dt/2);


		p4x = p1x + v3x * (dt);
		p4y = p1y + v3y * (dt);
		var force = this.force({x:p4x, y:p4y});
		var fx = force.x;//-1000 * 1.85 * distance.x / Math.pow(distanceAbs,3);
		var fy = force.y;//-1000 * 1.85 * distance.y / Math.pow(distanceAbs,3);
		a4x = fx;
		a4y = fy;

		v4x = v1x + a3x * (dt);
		v4y = v1y + a3y * (dt);






		this.x = this.x + (v1x + 2*v2x + 2*v3x + v4x)*dt/6;
		this.y = this.y + (v1y + 2*v2y + 2*v3y + v4y)*dt/6;

		this.vx = this.vx + (a1x + 2*a2x + 2*a3x + a4x)*dt/6;
		this.vy = this.vy + (a1y + 2*a2y + 2*a3y + a4y)*dt/6;
	*/

	position := p1.AddVec(v1.AddVec(v2.Multiply(2.0)).AddVec(v3.Multiply(2.0)).AddVec(v4).Multiply(delta / 6.0))

	planet.rigidbody.Position = position

	velocity := v1.AddVec(force1.AddVec(force2.Multiply(2.0)).AddVec(force3.Multiply(2.0)).AddVec(force4).Multiply(delta / 6.0))

	planet.rigidbody.Velocity = velocity

	log.Printf("position: %v", position)
	log.Printf("velocity: %v", velocity)
}

func (planet *Planet) Color() int64 {
	return planet.color
}

func (planet *Planet) Type() GameObjectType {
	return PlanetObject
}

func (planet *Planet) Write() *data.Planet {
	return &data.Planet{
		Color:    planet.color,
		Mass:     planet.rigidbody.Mass,
		Position: planet.initialPosition.Clone(),
		Radius:   planet.radius,
		Force:    planet.initialForce,
	}
}
