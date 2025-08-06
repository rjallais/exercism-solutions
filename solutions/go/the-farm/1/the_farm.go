package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

func (i *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %v", i.cows, i.message)
}

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	var (
		amount, err1 = fc.FodderAmount(cows)
		factor, err2 = fc.FatteningFactor()
		total        = amount * factor
	)
	if err1 != nil {
		return 0, err1
	}
	if err2 != nil {
		return 0, err2
	}
	return total / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fc, cows)
	} else {
		return 0, errors.New("invalid number of cows")
	}
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	default:
		return nil
	}
}
