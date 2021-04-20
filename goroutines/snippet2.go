package main

import "fmt"

func main() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	fmt.Println(counter)
}
