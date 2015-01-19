package main

import (
	"image/png"
	"os"
	"ragrace"
)

func main() {
	w := ragrace.NewWorld()
	c := new(ragrace.Camera)
	c.ResWidth = 512
	c.ResHeight = 512
	c.ViewPoint = ragrace.Vector{0.0, 0.0, -5.0}
	c.Depth = 5.0
	c.Width = 2.0
	c.Height = 2.0
	w.SetCamera(c)

	s := &ragrace.Sphere{R: 1.0}
	l := ragrace.NewLocate(s)
	l.Loc = ragrace.Vector{0.0, 0.0, 5.0}
	w.SetObject(l)

	img, err := w.Rendering()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	file, err := os.Create("example.png")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	if err := png.Encode(file, img); err != nil {
		panic(err)
		os.Exit(1)
	}
}
