package ragrace

import "math"

type ShapeSphere struct {
	Shape
	R float64
}

func (s *ShapeSphere) Collision(ray Ray) float64 {
	a := ray.D.Dot(Vector(ray.D))
	b := ray.S.Dot(Vector(ray.D)) * 2.0
	c := ray.S.Dot(ray.S) - s.R*s.R
	d := b*b - 4*a*c
	if d < 0.0 {
		return fNaN
	}
	t1 := (-b + math.Sqrt(d)) / (2 * a)
	t2 := (-b - math.Sqrt(d)) / (2 * a)
	if t1 < 0.0 {
		if t2 < 0.0 {
			return fNaN
		}
		return t1
	}
	if t2 < 0.0 {
		return t2
	}
	return math.Min(t1, t2)
}

func (s *ShapeSphere) Visual(v Vector) Material {
	return DefaultMaterial
}
