package task1

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"both zeros", 0, 0, 0},
		{"positive and zero", 5, 0, 5},

		{"negative numbers", -2, -3, -5},
		{"negative and positive", -5, 10, 5},
		{"positive and negative", 7, -3, 4},

		{"large numbers", 1000, 2000, 3000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)

			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d, expexted %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		expected int
	}{
		{"positive numbers", 10, 5, 5},
		{"both zeros", 0, 0, 0},
		{"positive and zero", 10, 0, 10},

		{"negative numbers", -5, -3, -2},
		{"negative and positive", -5, 2, -7},
		{"positibe and negative", 5, -3, 8},

		{"large numbers", 3000, 1000, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			result := Subtract(tt.a, tt.b)

			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		expected int
	}{
		{"positive numbers", 2, 5, 10},
		{"negative numbers", -5, -2, 10},
		{"both zeros", 0, 0, 0},
		{"positive and zero", 4, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			result := Multiply(tt.a, tt.b)

			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDevide(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int 
		wantErr error
	}{
		{"normal division", 10, 5, 2, nil},
		{"division by zero", 10, 0, 0, ErrDivisionByZero},
	}
	for _, tt := range tests {
		t.Run(tt. name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Divide(%d, %d) error = %v, want %v", tt.a, tt.b, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
