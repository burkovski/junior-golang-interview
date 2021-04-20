package main

import "fmt"

func main() {
	var ch chan int
	go func() {
		ch <- 5
	}()
	fmt.Println(<-ch)
}
