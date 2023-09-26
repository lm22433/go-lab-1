package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

// The double function takes a pointer to the slice.
// This is because we need to update this pointer with
// the newly sized slice that has been constructed by the
// concatenation of the two slices.
func double(slice *[]int) {
	*slice = append(*slice, *slice...)
}

// Go passes slices to functions by reference, so we have direct access
// to the slice itself and the underlying array. This means that we can
// easily update values of a slice.
func mapSlice(f func(a int) int, slice []int) {
	for i, x := range slice {
		slice[i] = f(x)
	}
}

// Go passes arrays to functions by value which means that a copy is
// made. Any changes made to the copy do not affect the original array.
// As a result we need to pass a pointer to the original array to the function
func mapArray(f func(a int) int, array *[3]int) {
	for i, x := range array {
		array[i] = f(x)
	}
}

func main() {
	// A slice is a reference to a continuous segment of am array.
	// It is a reference type. (Passed by reference)
	intsSlice := []int{1, 2, 3}

	// An array is a fixed length, homogeneous data structure.
	// It is a value type. (Passed by value)
	intsArray := [3]int{1, 2, 3}

	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	// We pass the array as a pointer so that we have access to the
	// underlying memory so that the values can be directly modified.
	mapArray(addOne, &intsArray)
	fmt.Println(intsArray)

	// Here is interesting because the underlying array behind the slice
	// is changed. As a result, both slices are changed.
	newSlice := intsSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println(intsSlice, newSlice)

	// Here we pass the slice as a reference so that when the slice
	// is appended to itself we can update the reference to that slice
	// to reflect that the size of the slice has changed.
	double(&intsSlice)
	fmt.Println(intsSlice)
}
