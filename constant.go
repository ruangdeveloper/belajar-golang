package main

import "fmt"

func main(){
    const name string = "Ruang Developer"
    // Tidak wajib menyebutkan tipe data
    const web = "www.ruangdeveloper.com"

	// Deklarasi multiple constant
	const (
		series string = "Belajar Golang Dari Dasar"
		vote = 100
	)

	fmt.Println(name)
	fmt.Println(web)
	fmt.Println(series)
	fmt.Println(vote)
}