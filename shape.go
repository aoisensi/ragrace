package ragrace

type Shape interface {
	Object
	Collision(Ray) float64
	Visual(Vector) Material
}
