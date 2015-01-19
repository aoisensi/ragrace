package ragrace

import (
	"image"
	"image/color"
)

type World struct {
	obj Object
	cam *Camera
}

func NewWorld() *World {
	w := new(World)
	return w
}

func (w *World) SetCamera(camera *Camera) {
	w.cam = camera
}

func (w *World) SetObject(object Object) {
	w.obj = object
}

func (w *World) Rendering() (image.Image, error) {
	cam := w.cam
	img := image.NewRGBA(image.Rect(0, 0, cam.ResWidth, cam.ResHeight))
	xs := makeGridSlice(cam.ResWidth, cam.Width)
	ys := makeGridSlice(cam.ResHeight, cam.Height)
	for px, x := range xs {
		for py, y := range ys {
			to := Vector{x, y, cam.Depth}.Normalizing()
			rr := w.obj.Collision(Ray{cam.ViewPoint, to})
			if rr >= 0.0 {
				img.SetRGBA(px, py, color.RGBA{255, 0, 0, 255})
			}
		}
	}

	return img, nil
}

func makeGridSlice(res int, width float64) []float64 {
	r := make([]float64, res)
	resf := 1.0 / float64(res-1)
	for i := range r {
		r[i] = width*float64(i)*resf - (width / 2.0)
	}
	return r
}
