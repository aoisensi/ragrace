package ragrace

type ShapePlane struct {
	Shape
	N NVector
}

func NewPlane(n NVector) *ShapePlane {
	p := new(ShapePlane)
	p.N = n
	return p
}

func (p *ShapePlane) Collision(ray Ray) float64 {
	m := ray.D.Dot(Vector(p.N))
	if m > 0.0 {
		return fNaN
	}
	return ray.S.Dot(Vector(p.N)) / m
}

func (s *ShapePlane) Visual(v Vector) Material {
	return DefaultMaterial
}
