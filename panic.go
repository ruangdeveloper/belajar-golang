package main

import "fmt"

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	// kode ini akan tetap dieksekusi
	defer hello()
	
	// meskipun kita memanggil panic function
	panic("Panic dipanggil")
	
	// sedangkan kode ini tidak akan dieksekusi, karena program berhenti
	// setelah panic function dipanggil
	fmt.Println("Tidak akan dieksekusi")
}