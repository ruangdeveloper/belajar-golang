package main

import "fmt"

func main() {
	var x = 10
	var y = 5

	// Operasi penjumlahan
	var sum = x + y
	fmt.Println(sum)

	// Operasi pengurangan
	var sub = x - y
	fmt.Println(sub)

	// Operasi perkalian
	var mul = x * y
	fmt.Println(mul)

	// Operasi pembagian
	var div = x / y
	fmt.Println(div)

	// Operasi modulus (sisa pembagian)
	var mod = x % y
	fmt.Println(mod)

	// Augmented assignments
	x += 5
	fmt.Println(x)

	// Unary operator
	var positiveX = +x
	fmt.Println(positiveX)

	var negativeX = -x
	fmt.Println(negativeX)

	var boolean = true
	fmt.Println(boolean)

	var invertBoolean = !boolean
	fmt.Println(invertBoolean)
}
