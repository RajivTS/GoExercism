package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	CowCount int
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.CowCount)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	errNegativeFodder := errors.New("Negative fodder")
	errDivisonByZero := errors.New("Division by zero")
	errSillyNephew := &SillyNephewError{CowCount: cows}
	if fodderAmount < 0 {
		return 0, errNegativeFodder
	} else if cows == 0 {
		return 0, errDivisonByZero
	} else if cows < 0 {
		return 0, errSillyNephew
	} else if err == ErrScaleMalfunction {
		return fodderAmount * 2 / float64(cows), nil
	} else if err != nil {
		return 0, err
	} else {
		return fodderAmount / float64(cows), nil
	}
}
