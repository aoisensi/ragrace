package ragrace

type Sphere struct {
	Object
	R float64
}

func (s *Sphere) Collision(ray Ray) float64 {
	a := ray.D.Dot(Vector(ray.D))
	b := ray.S.Dot(Vector(ray.D)) * 2.0
	c := ray.S.Dot(ray.S) - s.R*s.R
	//fmt.Printf("%v", ray)
	return b*b - 4*a*c
}
