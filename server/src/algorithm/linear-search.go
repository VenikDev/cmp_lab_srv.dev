package algorithm

// LinearSearch
// With the `any` keyword, we define a type parameter `T` that can be any type,
// and use it to specify the types of the `arr` slice and the `pred` function argument.
// The `pred` argument takes a value of type `T` as input and returns a boolean value.
// The rest of the function remains the same,
// with a loop that iterates over each element of the `arr` slice and returns the index of the first element for which
// the `pred` function returns `true`, or `-1` if no such element is found.
// In the `main` function, we can now call the `LinearSearch` function with any type of slice and predicate function
// that satisfy its type constraints. Here,
// we demonstrate its usage by searching for an `int` value in an `[]int` slice,
// but this same function can be used with any type that satisfies the `any` constraint defined in the function
// signature.
func LinearSearch[T any](arr []T, pred func(T) bool) int {
	for i := 0; i < len(arr); i++ {
		if pred(arr[i]) {
			return i
		}
	}
	return -1
}
