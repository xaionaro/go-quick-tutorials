package main

import "fmt"

func main() {
	// This's an array:
	var ar [2]int

	ar[0] = 3
	ar[1] = 1416

	// This's a slice:
	sl := make([]int, 2)
	sl[0] = ar[1]
	sl[1] = ar[0]
	// NOTE: you cannot directly use slice as an array and vice versa (it won't compile),
	// but you can convert array to slice this way:
	sl2 := ar[:]
	sl2  = sl

	// adding an element to the slice:
	sl = append(sl, -314)

	// This's a map:
	m := make(map[string]string)
	m["ping"] = "pong"

	// Printing everything
	fmt.Printf("ar  == %v\nsl  == %v\nsl2 == %v\nm   == %v\nm[\"ping\"] == %v\n", ar, sl, sl2, m, m["ping"])
}
