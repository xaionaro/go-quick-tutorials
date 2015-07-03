package main

import "fmt"

func main() {
	// Learning assignments:
	var i0, i1 int
	i0  = 1		// i0 == 1
	i1  = i0	// i1 == 1
	i2 := i0	// i2 == 1; "i2" has the same type as "i0"
	i2 += 1		// i2 == 2

	var b int64

	b = int64(i0)	//  b == 1

	// Printing "i1", "i2"
	fmt.Printf("%v\n%v\n%v\n", i1, i2, b)
}
