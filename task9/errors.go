package task9

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrNegativeNumber = errors.New("negative number not allowed")
	ErrDivByZero      = errors.New("division by zero")
	ErrTooLarge       = errors.New("number too large")
)

func ParsePositiveNumber(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("%w: empty string", ErrInvalidInput)
	}

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	if num < 0 {
		return 0, fmt.Errorf("%w: %d", ErrNegativeNumber, num)
	}

	if num > 1000 {
		return 0, fmt.Errorf("%w: %d > 1000", ErrTooLarge, num)
	}

	return num, nil
}

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

func ValidateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("%w: age cannot be negative", ErrInvalidInput)
	}

	if age < 18 {
		return errors.New("must be at least 18 years old")
	}

	if age > 120 {
		return fmt.Errorf("%w: %d is not a valid age", ErrTooLarge, age)
	}

	return nil
}

