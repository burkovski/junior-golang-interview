package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[:]
	s2 := s[1:3]

	fmt.Printf("%#v\n", s)
	s1[0] = 42
	s2[0] = 42
	fmt.Printf("%#v\n", s)
}
