package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:3]

	fmt.Printf("%#v\n", s)
	appendToSlice(s)
	appendToSlice(s1)
	fmt.Printf("%#v\n", s)
}

func appendToSlice(dst []int) {
	dst = append(dst, 42)
}
