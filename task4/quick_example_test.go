package task4

import (
	"math"
	"strings"
	"testing"
	"testing/quick"
)

func TestIsEvenProperty(t *testing.T) {
	property := func(x int) bool {
		return IsEven(x) == IsEven(x+2)
	}
	if err := quick.Check(property, nil); err != nil {
		t.Errorf("IsEven property failed: %v", err)
	}
}

func TestCommutativeAdd(t *testing.T) {
	property := func(a, b int) bool {
		return CommutativeAdd(a, b)
	}
	if err := quick.Check(property, nil); err != nil {
		t.Errorf("CommutativeAdd failed: %v", err)
	}
}

func TestAssosiativeAdd(t *testing.T) {
	property := func(a, b, c int) bool {
		return AssosiativeAdd(a, b, c)
	}
	if err := quick.Check(property, nil); err != nil {
		t.Errorf("AssosiativeAdd failed: %v", err)
	}
}

func TestAdditiveInverse(t *testing.T) {
	property := func(n int) bool {
		return AdditiveInverse(n)
	}

	config := &quick.Config{
		MaxCount:      1000,
		MaxCountScale: 1.0,
	}

	if err := quick.Check(property, config); err != nil {
		t.Errorf("AdditiveInverse failed: %v", err)
	}
}

func TestPalindromeProperty(t *testing.T) {
	property := func(s string) bool {
		normalized := strings.ToLower(strings.ReplaceAll(s, " ", ""))
		if IsPalindrome(s) {
			return Reverse(normalized) == normalized
		}
		return true
	}

	config := &quick.Config{
		MaxCount: 500,
	}

	if err := quick.Check(property, config); err != nil {
		t.Errorf("IsPalindrome fauled: %v", err)
	}
}

func TestHasPositiveSum(t *testing.T) {
	tests := []struct {
        name  string
        nums  []int
        want  bool
    }{
        {"пустой слайс", []int{}, false},
        {"только отрицательные", []int{-1, -2, -3}, false},
        {"только нули", []int{0, 0, 0}, false},
        {"одно положительное", []int{5}, true},
        {"несколько положительных", []int{1, 2, 3}, true},
        {"смешанные", []int{-5, 10, -3, 2}, true},
        {"большие числа", []int{math.MaxInt64, 1}, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := HasPositiveSum(tt.nums)
            if got != tt.want {
                t.Errorf("HasPositiveSum(%v) = %v, want %v", tt.nums, got, tt.want)
            }
        })
	}
}

// Интеграционный тест с quick
func TestAllProperties(t *testing.T) {
	properties := []struct {
		name     string
		property interface{}
	}{
		{"CommutativeAdd", func(a, b int) bool { return CommutativeAdd(a, b) }},
		{"AssosiativeAdd", func(a, b, c int) bool { return AssosiativeAdd(a, b, c) }},
	}

	config := &quick.Config{
		MaxCount: 100,
	}

	for _, p := range properties {
		t.Run(p.name, func(t *testing.T) {
			if err := quick.Check(p.property, config); err != nil {
				t.Errorf("%s failed: %v", p.name, err)
			}
		})
	}
}
