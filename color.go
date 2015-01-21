package ragrace

type Color interface {
	RGB() (r, g, b float64)
}

func ColorMul(l, r Color) Color {
	ar, ag, ab := l.RGB()
	br, bg, bb := r.RGB()
	return NewRGB(ar*br, ag*bg, ab*bb)
}

type RGB struct {
	Color
	R, G, B float64
}

func NewRGB(r, g, b float64) *RGB {
	c := new(RGB)
	c.R = r
	c.G = g
	c.B = b
	return c
}

func (c *RGB) RGB() (r, g, b float64) {
	return c.R, c.G, c.B
}

type Gray struct {
	Color
	G float64
}

func NewGray(g float64) *Gray {
	c := new(Gray)
	c.G = g
	return c
}

func (c *Gray) RGB() (r, g, b float64) {
	return c.G, c.G, c.G
}
