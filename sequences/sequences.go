package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice *[]int) {
	*slice = append(*slice, *slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, x := range slice {
		slice[i] = f(x)
	}
}

func mapArray(f func(a int) int, array *[3]int) {
	for i, x := range array {
		array[i] = f(x)
	}
}

func main() {
	// A slice is a reference to a continuous segment of am array.
	// It is a reference type. (Passed by reference)
	intsSlice := []int{1, 2, 3}

	// An array is a fixed length, homogeneouus data structure.
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
