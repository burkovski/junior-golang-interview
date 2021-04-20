package main

import "fmt"

func main() {
	var s []string
	s = append(s, s...)
	s = append(s, "foobar")
	s = append(s, s...)

	for i := range s {
		fmt.Printf("%#v\n", s[i])
	}
}
