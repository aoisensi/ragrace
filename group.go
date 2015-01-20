package ragrace

import "math"

type Group struct {
	Object
	objs []Object
}

func NewGroup(size int) *Group {
	g := new(Group)
	g.objs = make([]Object, size)
	return g
}

func NewGroupSlice(objs []Object) *Group {
	g := new(Group)
	g.objs = objs
	return g
}

func (g *Group) Collision(ray Ray) float64 {
	r := fNaN
	for _, o := range g.objs {
		if o == nil {
			continue
		}
		s := o.Collision(ray)
		if isNaN(r) {
			r = s
			continue
		}
		if !isNaN(s) {
			r = math.Min(r, s)
		}
	}

	return r
}
