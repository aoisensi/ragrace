package ragrace

import "image"

type World struct {
	cam *Camera
}

func NewWorld() *World {
	w := new(World)
	return w
}

func (w *World) SetCamera(camera *Camera) {
	w.cam = camera
}

func (w *World) Rendering() (image.Image, error) {
	img := image.NewRGBA(w.cam.size)
	return image.Image(img), nil
}
