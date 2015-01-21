package ragrace

var (
	DefaultMaterial = Material{
		ARC:       0.01,
		DRC:       0.69,
		SRC:       0.30,
		Shininess: 8,
		Color:     &Gray{G: 1.0},
	}
)

type Material struct {
	ARC, DRC, SRC float64
	Shininess     float64
	Color         Color
}
