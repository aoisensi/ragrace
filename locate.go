package ragrace

type Locate struct {
	Shape
	Obj Shape
	Loc Vector
}

func NewLocate(obj Shape) *Locate {
	l := new(Locate)
	l.Obj = obj
	return l
}

func (l *Locate) Collision(ray Ray) float64 {
	return l.Obj.Collision(ray.Move(l.Loc.Inverse()))
}
