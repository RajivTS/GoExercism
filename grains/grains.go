package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	errOutOfRange := errors.New("invalid square number")
	if number <= 0 || number > 64 {
		return 0, errOutOfRange
	} else {
		return uint64(math.Pow(2, float64(number-1))), nil
	}
}

func Total() uint64 {
	squareVal, _ := Square(64)
	return 2*squareVal - 1
}
