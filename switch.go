package main

import "fmt"

func main() {
	var category = "CAT1"

	switch category {
	case "CAT1":
		fmt.Println("Kategori CAT1")
	case "CAT2":
		fmt.Println("Kategori CAT2")
	default:
		fmt.Println("Kategori tidak dikenali")
	}

	switch married := true; married {
	case true:
		fmt.Println("Sudah menikah")
	case false:
		fmt.Println("Belum menikah")
	}

	id := "A1"
	switch {
	case id == "A1":
		fmt.Println("ID = A1")
	case id == "A2":
		fmt.Println("ID = A2")
	default:
		fmt.Println("ID Invalid")
	}
}
