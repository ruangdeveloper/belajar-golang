package main

import "fmt"

func checkSudahMakan(sudahMakan bool) {
	if sudahMakan == false {
		panic("Ups, kamu belum makan")
	}

	fmt.Println("Mantap, sudah makan")
}

func selesai() {
	er := recover()

	if er != nil {
		fmt.Println("Ada eror:", er)
	}

	fmt.Println("Selesai")
}

func main() {
	defer selesai()
	checkSudahMakan(false)
}
