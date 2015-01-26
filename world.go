package ragrace

import (
	"image"
	"image/color"
)

type World struct {
	obj    Shape
	cam    *Camera
	lights []Light
}

func NewWorld() *World {
	w := new(World)
	return w
}

func (w *World) SetCamera(camera *Camera) {
	w.cam = camera
}

func (w *World) SetObject(object Shape) {
	w.obj = object
}

func (w *World) Rendering() (image.Image, error) {
	cam := w.cam
	img := image.NewRGBA(image.Rect(0, 0, cam.ResWidth, cam.ResHeight))
	xs := makeGridSlice(cam.ResWidth, cam.Width)
	ys := makeGridSlice(cam.ResHeight, -cam.Height)
	for px, x := range xs {
		for py, y := range ys {
			to := Vector{x, y, cam.Depth}.Normalizing()
			img.Set(px, py, w.emmitRay(Ray{cam.ViewPoint, to}))
		}
	}

	return img, nil
}

func (w *World) SetLights(l []Light) {
	w.lights = l
}

func (w *World) emmitLight(mat Material) Color {
	return NewGray(1.0)
}

func (w *World) emmitRay(ray Ray) color.Color {
	r := w.obj.Collision(ray)
	if isNaN(r) {
		return color.RGBA{}
	}
	return color.RGBA{255, 0, 0, 255}
}

func makeGridSlice(res int, width float64) []float64 {
	r := make([]float64, res)
	resf := 1.0 / float64(res-1)
	for i := range r {
		r[i] = width*float64(i)*resf - (width / 2.0)
	}
	return r
}
