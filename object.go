package ragrace

type Object interface {
	Collision(ray Ray) float64
}
