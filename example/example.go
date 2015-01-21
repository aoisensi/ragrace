package main

import (
	"image/png"
	"os"

	rg ".."
)

func main() {
	w := rg.NewWorld()
	c := new(rg.Camera)
	c.ResWidth = 512
	c.ResHeight = 512
	c.ViewPoint = rg.Vector{0.0, 0.0, -5.0}
	c.Depth = 5.0
	c.Width = 2.0
	c.Height = 2.0
	w.SetCamera(c)

	s := &rg.ShapeSphere{R: 1.0}
	sl := rg.NewLocate(s)
	sl.Loc = rg.Vector{0.0, 0.0, 5.0}

	p := rg.NewPlane(rg.NVector{0.0, 1.0, 0.0})
	pl := rg.NewLocate(p)
	pl.Loc = rg.Vector{0.0, -1.0, 0.0}

	w.SetObject(rg.NewGroupSlice([]rg.Object{sl, pl}))

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
