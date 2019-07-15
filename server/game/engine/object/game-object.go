package object

type GameObject interface {
	ID() int64
	Radius() float64
	RigidBody() *RigidBody
}
