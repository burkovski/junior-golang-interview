package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()

	_, _ = <-worker(), <-worker()

	elapsed := time.Since(timeStart).Seconds()
	fmt.Println(int(elapsed))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()
	return ch
}
