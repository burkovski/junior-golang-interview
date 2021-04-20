package main

import "fmt"

type User struct {
	ID         int
	IsVerified bool
}

func main() {
	users := []User{{1, true}, {2, false}, {3, true}}
	verified := filterVerifiedUsers(users)
	for _, user := range verified {
		fmt.Println(user.ID)
	}
}

func filterVerifiedUsers(users []User) []*User {
	var verified []*User
	for _, user := range users {
		if user.IsVerified {
			verified = append(verified, &user)
		}
	}
	return verified
}
