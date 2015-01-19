package ragrace

type Camera struct {
	ResWidth, ResHeight  int
	Width, Height, Depth float64
	ViewPoint            Vector
}

func NewCamera() *Camera {
	return new(Camera)
}
