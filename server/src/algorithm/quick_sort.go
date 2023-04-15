package algorithm

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
