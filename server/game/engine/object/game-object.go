package object

type GameObject interface {
	ID() int64
	Body() *Body
}
