package world

type World struct {
	tick    int64
	objects map[int64]Object
}

func NewWorld() *World {
	return &World{
		tick:    0,
		objects: make(map[int64]Object),
	}
}

func (world *World) GetObject(id int64) Object {
	object, ok := world.objects[id]
	if !ok {
		return nil
	}
	return object
}

func (world *World) AddObject(object Object) {
	id := object.ID()
	world.objects[id] = object
}

func (world *World) RemoveObject(id int64) Object {
	object, ok := world.objects[id]
	if !ok {
		return nil
	}

	delete(world.objects, id)

	return object
}

func (world *World) GetGameObjectsInSphere(target Object, radius float64) map[int64]Object {
	// get the rigidbody
	targetRigidbody := target.RigidBody()
	// get the radius
	targetRadius := target.Radius()

	contains := make(map[int64]Object)
	for id, object := range world.objects {
		// get the body of the game object
		objectRigidbody := object.RigidBody()

		// clauclate the distance
		distance := targetRigidbody.Position.DistanceTo(objectRigidbody.Position)

		if (distance - targetRadius) <= radius {
			contains[id] = object
		}
	}

	return contains
}

func (world *World) GetTick() int64 {
	return world.tick
}

func (world *World) GetObjects() map[int64]Object {
	return world.objects
}

func (world *World) Update(delta float64) {
	world.tick++
	for _, gameObject := range world.objects {
		gameObject.Update(delta)
	}
}
