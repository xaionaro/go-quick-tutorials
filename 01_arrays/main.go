package main

import "fmt"

func main() {
	// Learning assignments:
	var i0, i1 int
	i0  = 1
	i1  = i0
	i2 := i0	// "i2" will have the same type as "i0"

	// This's an array:
	var ar [2]int

	ar[0] = 3
	ar[1] = 1416

	// This's a slice:
	sl := make([]int, 2);
	sl[0] = ar[0]
	sl[1] = ar[1]

	// NOTE: you cannot directly use slice as an array and vice versa (it won't compile)

	// Printing "i1", "i2", "ar" and "sl"
	fmt.Printf("%v\n%v\n%v\n%v\n", i1, i2, ar, sl);
}
