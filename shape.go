package ragrace

type Shape interface {
	Collision(Ray) float64
	Visual(Vector) Material
}
