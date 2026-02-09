package task3

import (
	"reflect"
	"slices"
	"testing"
)

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "обычный случай",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "пустой срез",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "только нечетные",
			input:    []int{1, 3, 5},
			expected: []int{},
		},
		{
			name:     "только четные",
			input:    []int{2, 4, 6},
			expected: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterEven(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FilterEven(%v) = %v, ожидали %v", tt.input, got, tt.expected)
			}

			if !slices.Equal(got, tt.expected) {
				t.Logf("(альтернативная проверка) слайсы не равны")
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "уникальные строки",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "с дубликатами",
			input:    []string{"a", "b", "c", "a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "пустой слайс",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "все одинаковые",
			input:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.input)

			if !slices.Equal(got, tt.expected) {
				t.Errorf("Unique(%v) = %v, ожидали %v", tt.input, got, tt.expected)
			}

			seen := make(map[string]bool)
			for _, item := range got {
				if seen[item] {
					t.Errorf("Найден дубликат в рузельтате: %v", item)
				}
				seen[item] = true
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]int
	}{
		{
			name:  "обычный случай",
			input: []string{"apple", "banana", "apple", "orange", "banana", "apple"},
			expected: map[string]int{
				"apple":  3,
				"banana": 2,
				"orange": 1,
			},
		},
		{
			name:     "пустой слайс",
			input:    []string{},
			expected: map[string]int{},
		},
		{
			name:  "одно слово",
			input: []string{"hello", "hello", "hello"},
			expected: map[string]int{
				"hello": 3,
			},
		},
		{
			name:  "регистр имеет значение",
			input: []string{"Go", "go", "GO"},
			expected: map[string]int{
				"Go": 1,
				"go": 1,
				"GO": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.input)

			// Для мапов используем reflect.DeepEqual
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CountWords(%v) = %v, ожидали %v",
					tt.input, got, tt.expected)
			}
			totalCount := 0
			for _, count := range got {
				totalCount += count
			}

			if totalCount != len(tt.input) {
				t.Errorf("Общее количество слов не совпадает: %d != %d",
					totalCount, len(tt.input))
			}
		})
	}
}
