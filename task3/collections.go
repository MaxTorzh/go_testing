package task3

func FilterEven(numbers []int) []int {
	result := []int{}
	for _, n := range numbers {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

func Unique(strings []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, s := range strings {
		if !seen[s] {
			seen[s] = true
			result = append(result, s)
		}
	}
	return result
}

func CountWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
