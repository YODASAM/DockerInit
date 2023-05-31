package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	// fmt.Println("go" + "lang")

	// fmt.Println("1+1 =", 1+1)
	// fmt.Println("7.0/3.0 =", 7.0/3.0)

	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	var t = "testing"
	fmt.Println(t)
	fmt.Println(t, "------>", s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println("int64 version of d : ", int64(d))

	const ck = 27365
	fmt.Println(ck + 23 - 723654)
	fmt.Println("Sin(ck)=", math.Sin(ck))

	var b, c int = 1, 2
	fmt.Println(b, c)

	var ad = true
	fmt.Println(ad)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

}
