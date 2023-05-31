package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define a slice of integers.
	numbers := []int{5, 1, 4, 2, 8, 3}

	fmt.Println("Before sorting: ", numbers)

	// Use sort.Ints to sort the slice.
	sort.Ints(numbers)

	fmt.Println("After sorting: ", numbers)
}
