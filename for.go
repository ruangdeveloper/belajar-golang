package main

import "fmt"

func main() {
	count := 1
	for count <= 5 {
		fmt.Println("Perulangan:", count)
		count++
	}

	for count := 1; count <= 5; count++ {
		fmt.Println("Perulangan:", count)
	}

	fruits := []string{"Nanas", "Melon", "Semangka"}

	for index, fruit := range fruits {
		fmt.Println("Index:", index, "=", fruit)
	}

	fmt.Println("---")

	website := map[string]string{
		"name":   "Ruang Developer",
		"domain": "www.ruangdeveloper.com",
	}

	for key, value := range website {
		fmt.Println(key, ":", value)
	}

	// mengganti variabel index menjadi underscore
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

}
