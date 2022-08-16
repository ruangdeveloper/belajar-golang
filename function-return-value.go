package main

import "fmt"

func sum(value1 float32, value2 float32) float32 {
	return value1 + value2
}

func isMarried(married bool) string {
	if married {
		return "Sudah nikah"
	} else {
		return "Belum nikah"
	}
}

func rectangle(sideLength float32) (float32, float32) {
	area := sideLength * sideLength
	perimeter := sideLength * 4

	return area, perimeter
}

func getHello() (hello string, world string) {
	hello = "Halo"
	world = "Dunia"

	return
}

func main() {
	result := sum(10, 10)

	fmt.Println("Hasil:", result)

	isMarried := isMarried(false)

	fmt.Println(isMarried)

	area, perimeter := rectangle(3)

	fmt.Println("Luas Persegi:", area)
	fmt.Println("Keliling Persegi:", perimeter)

	hello, world := getHello()
	fmt.Println(hello)
	fmt.Println(world)
}
