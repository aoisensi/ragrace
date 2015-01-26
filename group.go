package ragrace

import "math"

type Group struct {
	Shape
	objs []Shape
}

func NewGroup(size int) *Group {
	g := new(Group)
	g.objs = make([]Shape, size)
	return g
}

func NewGroupSlice(objs []Shape) *Group {
	g := new(Group)
	g.objs = objs
	return g
}

func (g *Group) Collision(ray Ray) float64 {
	r := fNaN
	for _, o := range g.objs {
		switch t := o.(type) {
		case Shape:
			s := t.Collision(ray)
			if isNaN(r) {
				r = s
				continue
			}
			if !isNaN(s) {
				r = math.Min(r, s)
			}
		}
	}

	return r
}
