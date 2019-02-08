package triangle

import "math"

// Kind represents return code for triangle
type Kind int

const (
	// NaT is not a triangle
	NaT = 0
	// Equ is an equilateral triangle
	Equ = 1
	// Iso is an isosceles triangle
	Iso = 2
	// Sca is a scalene triangle
	Sca = 3
)

// KindFromSides identifies the kind of a triangle based on shape side lengths
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if !isValidTriangle(a, b, c) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}

	return k
}

func isValidTriangle(a, b, c float64) bool {
	triangleInequality := (a+b >= c && a+c >= b && b+c >= a)
	return triangleInequality && isValidLength(a) && isValidLength(b) && isValidLength(c)
}

func isValidLength(number float64) bool {
	nan := math.NaN()
	pinf := math.Inf(1)
	ninf := math.Inf(-1)

	return number != nan && number != pinf && number != ninf && number > 0
}
