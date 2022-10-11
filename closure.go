package main

import "fmt"

func main() {
	name := "Rizky Kurniawan"

	printName := func() {
		// anggap saja tidak sengaja merubah isi variabel name
		name = "Tony Stark"
		fmt.Println("Di dalam function printName")
		fmt.Println(name)
	}
	
	printName()
	
	fmt.Println("Di luar function printName")
	fmt.Println(name)
}
