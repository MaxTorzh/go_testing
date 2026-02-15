package task4

import (
	"math/big"
	"strings"
)

// Математические свойства

func IsEven(n int) bool {
	return n%2 == 0
}

func CommutativeAdd(a, b int) bool {
	return (a + b) == (b + a)
}

func AssosiativeAdd(a, b, c int) bool {
	return (a+b)+c == a+(b+c)
}

func AdditiveInverse(n int) bool {
	return n+(-n) == 0
}

// Строковые свойства

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == Reverse(s)
}

// Слайсы

func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func HasPositiveSum(nums []int) bool {
    sum := big.NewInt(0)
    hasPositive := false
    
    for _, n := range nums {
        if n > 0 {
            hasPositive = true
            sum.Add(sum, big.NewInt(int64(n)))
        }
    }
    
    return hasPositive && sum.Sign() > 0
}
