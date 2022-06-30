package main

import "fmt"

func main() {
	var lolosWawancara = true
	var lolosKoding = true

	var diterimaKerja = lolosWawancara && lolosKoding

	fmt.Println(diterimaKerja)

	fmt.Println("And")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println("Or")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println("Not")
	fmt.Println(!true)
	fmt.Println(!false)
}
