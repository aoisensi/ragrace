package ragrace

type Ray struct {
	S Vector
	D NVector
}

func (r Ray) Move(vec Vector) Ray {
	return Ray{r.S.Add(vec), r.D}
}
