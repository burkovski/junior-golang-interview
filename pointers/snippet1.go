package main

import (
	"fmt"
)

func main() {
	v := 42
	fmt.Println(v)
	assignToPointer(&v)
	fmt.Println(v)
}

func assignToPointer(dst *int) {
	v := 13
	dst = &v
}
