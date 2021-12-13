package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError interface {
	customError(float64) error
}

func customError(cowNum int) error {
	return errors.New(fmt.Sprintf("silly nephew, there cannot be %d cows", cowNum))
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fod, err := weightFodder.FodderAmount()

	if err != nil {
		switch err {
		case ErrScaleMalfunction:
			fod *= 2
		default:
			return 0, err
		}
	}
	if fod < 0 {
		return 0, errors.New("Negative fodder")
	}
	if cows <= 0 {
		if cows == 0 {
			return 0, errors.New("Division by zero")
		} else {

			return 0, customError(cows)
		}
	}
	return fod / float64(cows), nil
	panic("Please implement DivideFood")
}
