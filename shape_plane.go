package ragrace

type Plane struct {
	Object
	N NVector
}

func NewPlane(n NVector) *Plane {
	p := new(Plane)
	p.N = n
	return p
}

func (p *Plane) Collision(ray Ray) float64 {
	m := ray.D.Dot(Vector(p.N))
	if m > 0.0 {
		return fNaN
	}
	return ray.S.Dot(Vector(p.N)) / m
}
