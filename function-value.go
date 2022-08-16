package main

import "fmt"

func sayHiTo(name string) string {
	return "Hi, " + name
}

func main() {
	hi := sayHiTo

	result := hi("Rizky")

	fmt.Println(result)
}
