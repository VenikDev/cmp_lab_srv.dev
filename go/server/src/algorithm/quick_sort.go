package algorithm

// QuickSort
// This is an implementation of the quicksort algorithm in Go,
// which sorts a slice of any type T by applying the "less" predicate on elements.
// The function takes two arguments:
// arr is the slice of elements to be sorted.
// less is a function that takes two elements of type T and returns a boolean value indicating whether the first
// element is less than the second element.
// If the length of the input slice is less than or equal to 1, it means the slice is already sorted,
// and so we return it as is. Otherwise,
// we select a pivot element from the middle of the slice and partition the rest of the elements into two slices: a
// left slice for elements that are less than the pivot,
// and a right slice for elements that are greater than or equal to the pivot.
// We then recursively sort the left and right slices using the same quicksort algorithm until they are sorted,
// and then merge them with the pivot element using the append() function to return the final sorted slice.
func QuickSort[T any](arr []T, less func(T, T) bool) []T {
	if len(arr) <= 1 {
		return arr
	}

	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	left := make([]T, 0, len(arr))
	right := make([]T, 0, len(arr))
	for i, val := range arr {
		if i == pivotIndex {
			continue
		}
		if less(val, pivot) {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	left = QuickSort(left, less)
	right = QuickSort(right, less)

	return append(append(left, pivot), right...)
}
