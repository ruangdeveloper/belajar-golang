package main

import "fmt"

func main(){
	var age1 int64 = 21
    // konversi ke integer 32
    var age2 int32 = int32(age1)

    fmt.Println(age1)
    fmt.Println(age2)
	
    var name = "Ruang Developer"
    var charR = name[0]
    var charRString = string(charR)

    fmt.Println(name)
    fmt.Println(charR)
    fmt.Println(charRString)
}