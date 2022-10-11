package main

import "fmt"

type User struct {
	Name, Email string
	Age         int
}

func main() {
	var user User

	user.Name = "Rizky"
	user.Email = "email@email.com"
	user.Age = 22

	fmt.Println(user)

	fmt.Println(user.Name)
	fmt.Println(user.Email)
	fmt.Println(user.Age)

	// struct literals
	var user1 = User{
		Name:  "User 1",
		Email: "email1@email.com",
		Age:   22,
	}

	fmt.Println(user1)

	var user2 = User{
		"User 2",
		"email2@email.com",
		20,
	}

	fmt.Println(user2)
}
