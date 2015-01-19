package ragrace

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Add(vec Vector) Vector {
	return Vector{
		v.X + vec.X,
		v.Y + vec.Y,
		v.Z + vec.Z,
	}
}

func (v Vector) Sub(vec Vector) Vector {
	return v.Add(vec.Inverse())
}

func (v Vector) Inverse() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}

func (v Vector) Mul(val float64) Vector {
	return Vector{
		v.X * val,
		v.Y * val,
		v.Z * val,
	}
}

func (v Vector) Norm() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v Vector) Normalize(norm float64) NVector {
	return NVector(v.Mul(1.0 / norm))
}

func (v Vector) Normalizing() NVector {
	return v.Normalize(v.Norm())
}

func (v Vector) Dot(vec Vector) float64 {
	return v.X*vec.X + v.Y*vec.Y + v.Z*vec.Z
}

func (v Vector) Cross(vec Vector) Vector {
	return Vector{
		v.Y*vec.Z - v.Z*vec.Y,
		v.Z*vec.X - v.X*vec.Z,
		v.X*vec.Y - v.Y*vec.X,
	}
}

//Normalized Vector
type NVector Vector

func (v NVector) Norm() float64 {
	return 1.0
}

func (v NVector) Normalize(norm float64) NVector {
	return v
}

func (v NVector) Normalizing() NVector {
	return v
}

func (v NVector) Dot(vec Vector) float64 {
	return Vector(v).Dot(Vector(vec))
}

func (v NVector) Cross(vec Vector) Vector {
	return Vector(v).Cross(Vector(vec))
}
