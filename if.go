package main

import "fmt"

func main() {
	var age = 22

	if age == 22 {
		fmt.Println(age)
	}

	var domain = "www.ruangdeveloper.com"

	if domain == "www.google.com" {
		fmt.Println("Kode ini dieksekusi karena kondisi benar")
	} else {
		fmt.Println("Kode ini dieksekusi karena kondisi salah")
	}

	var name = "Ruang Developer"

	if name == "Ruang Kelas" {
		fmt.Println("name = Ruang Kelas")
	} else if name == "Ruang Developer" {
		fmt.Println("name = Ruang Developer")
	} else {
		fmt.Println("semuanya salah")
	}

	if count := 100; count > 50 {
		fmt.Println("Count lebih dari 50")
	}
}
