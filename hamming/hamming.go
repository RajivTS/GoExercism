package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	lenA := utf8.RuneCountInString(a)
	lenB := utf8.RuneCountInString(b)
	var dist int
	if lenA != lenB {
		errMismatchLength := errors.New("The two strings are of different length")
		return dist, errMismatchLength
	} else {
		for i := 0; i < lenA; i++ {
			if a[i] != b[i] {
				dist++
			}
		}
		return dist, nil
	}
}
