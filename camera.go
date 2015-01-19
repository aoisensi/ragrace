package ragrace

import "image"

type Camera struct {
	size image.Rectangle
}

func (c *Camera) SetSize(width, height int) {
	c.size = image.Rectangle{
		Min: image.ZP,
		Max: image.Point{X: width, Y: height},
	}
}
