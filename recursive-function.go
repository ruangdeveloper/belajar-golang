package main

import "fmt"

func doFactorial(value int) int {
	if value == 1 {
		return 1
	}
	return value * doFactorial(value-1)
}

func doFactorialWithForLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	resultRecursive := doFactorial(5)
	fmt.Println(resultRecursive)

	resultForLoop := doFactorialWithForLoop(5)
	fmt.Println(resultForLoop)
}