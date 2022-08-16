package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Index:", i)
		if i == 3 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Index:", i)
	}
}
