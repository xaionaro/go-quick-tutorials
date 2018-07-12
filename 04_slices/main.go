package main

import "fmt"

func extend(s []int) { // won't work
	s = append(s, 10)

	m := make([]int, len(s), cap(s)*2)
	copy(m, s)
	s = m
}

func distort(s []int) { // will work
	s[0] = 9
}

func extendByPtr(s *[]int) { // will work, of course
	m := make([]int, len(*s), cap(*s)*2)
	copy(m, *s)
	*s = m
}

func main() {
	// see also example "02_arrays_and_maps"

	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("s:", s, len(s), cap(s))
	fmt.Println("s[2:4]:", s[2:4], len(s[2:4]), cap(s[2:4]))
	c := make([]int, len(s), cap(s)*2)
	copy(c, s)
	fmt.Println("c:", c, len(c), cap(c))
	m := append(s, 8)
	fmt.Println("m:", m, len(m), cap(m))
	m = append(m, c...)
	fmt.Println("m:", m, len(m), cap(m))
	extend(m)
	fmt.Println("m:", m, len(m), cap(m))
	distort(m)
	fmt.Println("m:", m, len(m), cap(m))
	extendByPtr(&m)
	fmt.Println("m:", m, len(m), cap(m))
}
