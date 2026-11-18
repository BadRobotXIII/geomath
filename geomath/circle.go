package geomath

//(c^2 = a^2 + b^2 - 2abcos(C))
//(b^2 = a^2 + c^2 - 2accos(B))
//(a^2 = b^2 + c^2 - 2bccos(A))

import (
	"math"
)

type Circle struct {
	Diameter     float64
	Radius       float64
	Circumfrence float64
}

// CalcCircum computes the circumfrence of a circle
//
// Parameters:
//
//	-r radius of circle
//
// returns the circumfrence of a radius
func (s Circle) CalcCircum(r float64) float64 {
	circum := 2 * math.Pi * r
	return circum
}

func (s Circle) Cord(a float64, b float64, c float64) float64 {
	c = a*a + b*b
	math.Sqrt(c)
	return c
}
