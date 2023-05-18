package tests

import (
	"comparisonLaboratories/src/algorithm"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{9, 2, 5, 1, 7, 6, 8, 3, 4}
	sorted := algorithm.QuickSort(arr, func(a, b int) bool {
		return a < b
	})
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !sliceEqual(sorted, expected) {
		t.Errorf("Expected %v but got %v", expected, sorted)
	}
}

func sliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, av := range a {
		if av != b[i] {
			return false
		}
	}
	return true
}
