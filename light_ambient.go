package ragrace

type LightAmbient struct {
	Light
	Color Color
}

func (l *LightAmbient) Spotlight(mat Material) Color {
	return ColorMul(l.Color, NewGray(mat.ARC))
}
