package triangle

import (
	"math"
	"sort"
)

// Kind is a type of triangle
type Kind int

const (
	// NaT is not a triangle
	NaT Kind = 0
	// Equ is equilateral
	Equ Kind = 1
	// Iso is isosceles
	Iso Kind = 2
	// Sca scalene
	Sca Kind = 3
)

// KindFromSides it return kind of triangle
func KindFromSides(a, b, c float64) Kind {

	s := []float64{a, b, c}

	for _, v := range s {
		if v <= 0 || math.IsNaN(v) || v > math.MaxFloat64 {
			return NaT
		}
	}

	sort.Float64s(s)

	if s[0]+s[1] < s[2] {
		return NaT
	}

	if s[0] == s[1] && s[0] == s[2] && s[1] == s[2] {
		return Equ
	}

	if s[0] == s[1] || s[0] == s[2] || s[1] == s[2] {
		return Iso
	}

	if s[0] != s[1] && s[0] != s[2] && s[1] != s[2] {
		return Sca
	}

	return NaT
}
