package ragrace

import "math"

var (
	fNaN  = math.NaN()
	fInfP = math.Inf(+1)
	fInfN = math.Inf(-1)
)

func isNaN(x float64) bool {
	return math.IsNaN(x)
}

func isInfP(x float64) bool {
	return math.IsInf(x, 1)
}
