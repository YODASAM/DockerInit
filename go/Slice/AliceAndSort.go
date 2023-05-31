/* package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define a slice of integers.
	numbers := []int{5, 1, 4, 8, 3}

	fmt.Println("Before sorting: ", numbers)

	// Use sort.Ints to sort the slice.
	sort.Ints(numbers)

	fmt.Println("After sorting: ", numbers)
}
*/

package main

import (
	"fmt"
	"reflect"
	//"log"
	// "net/http"
	//  "os"
)

func main() {
	// // Define a slice of integers.
	// numbers := []int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)

	// // Append to a slice.
	// numbers = append(numbers, 6)
	// fmt.Println(numbers)

	// // Slice a slice.
	// slice := numbers[1:4]
	// fmt.Println(slice)

	// // Change an element in the slice.
	// numbers[0] = 100
	// fmt.Println(numbers)
	//////fmt.Println(quote.Go())

	// // Go Slice
	// //  Make a slice with make function.
	// moreNumbers := make([]int, 3)
	// moreNumbers[0] = 7
	// moreNumbers[1] = 8
	// moreNumbers[2] = 9
	// fmt.Println(moreNumbers)

	// // Go Array
	// var newArray [4]int
	// newArray[0] = 76
	// g := newArray[0]
	// newArray[0] = newArray[0] + 23
	// fmt.Println(g + 7)
	// fmt.Println(newArray[0])

	// Adding string to a byte string  KEY : Use SUFFIX : "..."
	a := []byte("hello")
	s := "world"
	a = append(a, s...) // use "..." as suffice
	fmt.Printf("%s", a)

	fmt.Println("Go first example")

	// var i int = 1
	var w float64 = 12.5

	// fmt.Println(i, w)

	//REFLECT Example : dont need to import reflect package latest

	// var name = "John Doe"
	// var age = 34

	// fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(w))

	// fmt.Printf("%s is %d years old\n", name, age)

}
