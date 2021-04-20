package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			ch <- (i + 1) * 2
			wg.Done()
		}(i)
	}
	fmt.Println(<-ch)
	wg.Wait()
}
