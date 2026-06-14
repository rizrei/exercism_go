package thefarm

import (
	"errors"
	"fmt"
)

type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}

func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	totalAmountOffodder, err := fc.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalAmountOffodder / float64(numberOfCows) * fatteningFactor, nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fc, numberOfCows)
}

type InvalidCowsError struct {
	NumberOfCows int
	Message      string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{NumberOfCows: numberOfCows, Message: "there are no negative cows"}
	}
	if numberOfCows == 0 {
		return &InvalidCowsError{NumberOfCows: numberOfCows, Message: "no cows don't need food"}
	}

	return nil
}
