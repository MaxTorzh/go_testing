package task2

import (
	"errors"
	"fmt"
)

var (
	ErrNegativeNumber = errors.New("negative number not allowed")
	ErrZeroNumber     = errors.New("zero not allowed")
	ErrTooLarge       = errors.New("number too large")
)

func ValidateNumber(n int) error {
	if n < 0 {
		return ErrNegativeNumber
	}

	if n == 0 {
		return ErrZeroNumber
	}

	if n > 100 {
		return fmt.Errorf("%w: %d > 100", ErrTooLarge, n)
	}
	return nil
}

func ProcessValue(n int) (int, error) {
	if err := ValidateNumber(n); err != nil {
		return 0, fmt.Errorf("process failed: %w", err)
	}
	return n * 2, nil
}
