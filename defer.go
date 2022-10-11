package main

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	defer hello()
	
	div := 0
	result := 3/div

	fmt.Println(result)
}