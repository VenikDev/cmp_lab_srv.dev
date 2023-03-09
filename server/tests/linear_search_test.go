package tests

import (
	"comparisonLaboratories/src/algorithm"
	"testing"
)

func TestLinerSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	result := algorithm.LinearSearch(arr, func(x int) bool { return x == 3 })
	if result != 2 {
		t.Errorf("Expected result not found")
	}

	arrStr := []string{"apple", "orange", "banana"}
	result = algorithm.LinearSearch(arrStr, func(x string) bool { return x == "pear" })
	if result != -1 {
		t.Errorf("Expected result not found")
	}

	arrFloats := []float64{}
	result = algorithm.LinearSearch(arrFloats, func(x float64) bool { return x > 0 })
	if result != -1 {
		t.Errorf("Expected result not found")
	}

	arrFloats = []float64{1.5, 2.3, 3.7, 4.1, 5.2}
	result = algorithm.LinearSearch(arrFloats, func(f float64) bool {
		return f == 3.7
	})
	if result != 2 {
		t.Errorf("Expected index to be 2, but got %v", result)
	}

	arrFloats = []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	result = algorithm.LinearSearch(arrFloats, func(f float64) bool {
		return f >= 0.6
	})

	if result != -1 {
		t.Errorf("Expected index to be -1, but got %v", result)
	}

	type person struct {
		name string
		age  int
	}

	arrPersons := []person{
		{"John", 20},
		{"Jane", 25},
		{"Bob", 30},
	}

	result = algorithm.LinearSearch(arrPersons, func(p person) bool { return p.name == "Jane" })
	if result != 1 {
		t.Errorf("Expected result not found")
	}
}
