package task1

import (
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
