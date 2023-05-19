package tests

import (
	"cmp_lab/src/algorithm"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		pred     func(int) bool
		expected int
	}{
		{
			name:     "Search for an element that exists at the beginning of the int slice",
			arr:      []int{1, 2, 3},
			pred:     func(x int) bool { return x == 1 },
			expected: 0,
		},
		{
			name:     "Search for an element that exists at the end of the int slice",
			arr:      []int{1, 2, 3},
			pred:     func(x int) bool { return x == 3 },
			expected: 2,
		},
		{
			name:     "Search for an element that exists in the middle of the int slice",
			arr:      []int{1, 2, 3},
			pred:     func(x int) bool { return x == 2 },
			expected: 1,
		},
		{
			name:     "Search for an element that doesn't exist in the int slice",
			arr:      []int{1, 2, 3},
			pred:     func(x int) bool { return x == 4 },
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := algorithm.LinearSearch(tt.arr, tt.pred)
			if result != tt.expected {
				t.Errorf("Expected %d but got %d", tt.expected, result)
			}
		})
	}
}

func TestLinearSearchEmptySlice(t *testing.T) {
	result := algorithm.LinearSearch([]int{}, func(x int) bool { return true })
	if result != -1 {
		t.Errorf("Expected -1 but got %d", result)
	}
}

func TestLinearSearchNotFound(t *testing.T) {
	result := algorithm.LinearSearch([]string{"apple", "banana", "cherry"}, func(x string) bool { return x == "pear" })
	if result != -1 {
		t.Errorf("Expected -1 but got %d", result)
	}
}
