package main

import (
	"image/png"
	"os"
	"ragrace"
)

func main() {
	w := ragrace.NewWorld()
	c := new(ragrace.Camera)
	c.SetSize(1024, 1024)
	w.SetCamera(c)
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
