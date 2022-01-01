package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind uint

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

func GetMaxAndMin(a, b float64) (float64, float64) {
	if a >= b {
		return a, b
	} else {
		return b, a
	}
}
func GetSidesInOrder(a, b, c float64) (float64, float64, float64) {
	if a >= b && a >= c {
		b, c = GetMaxAndMin(b, c)
		return a, b, c
	} else if b >= a && b >= c {
		a, c = GetMaxAndMin(a, c)
		return b, a, c
	} else {
		a, b = GetMaxAndMin(a, b)
		return c, a, b
	}
}

// Implementation of the logic to decide the kind of triangle from the length of its sides
func KindFromSides(a, b, c float64) Kind {
	a, b, c = GetSidesInOrder(a, b, c)
	if a <= 0 || b <= 0 || c <= 0 || a > b+c {
		return NaT
	} else if a == b && b == c {
		return Equ
	} else if a == b || b == c || c == a {
		return Iso
	} else {
		return Sca
	}
}
