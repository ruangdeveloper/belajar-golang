package main

import "fmt"

func sumNumbers(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	result := sumNumbers(10, 10, 10)

	fmt.Println("Hasil:", result)
}
