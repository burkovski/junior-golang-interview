package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// Какой будет результат выполнения приложения
	wg := sync.WaitGroup{}
	data := make(map[string]int)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(value int) {
			k := strconv.Itoa(value)
			data[k] = value
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("%#v", data)
}
