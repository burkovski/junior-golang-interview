package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID int
}

func main() {
	users := make(map[int]User)
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			users[i] = User{i}
		}(i)
	}
	wg.Wait()
	fmt.Printf("%#v\n", users)
}
