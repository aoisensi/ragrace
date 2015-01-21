package ragrace

type Light interface {
	Spotlight(Material) Color
}
