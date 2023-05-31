package main

import "fmt"

func main() {
	// Define a slice of integers.
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// Append to a slice.
	numbers = append(numbers, 6)
	fmt.Println(numbers)

	// Slice a slice.
	slice := numbers[1:4]
	fmt.Println(slice)

	// Change an element in the slice.
	numbers[0] = 100
	fmt.Println(numbers)

	// Make a slice with make function.
	moreNumbers := make([]int, 3)
	moreNumbers[0] = 7
	moreNumbers[1] = 8
	moreNumbers[2] = 9
	fmt.Println(moreNumbers)
}
