package main

import "fmt"

// type User struct {
// 	ID int
// }

func main() {
	var users map[int]User
	for i := 0; i < 1000; i++ {
		users[i] = User{i}
	}
	fmt.Printf("%#v\n", users)
}
