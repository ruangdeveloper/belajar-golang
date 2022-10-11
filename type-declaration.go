package main

import "fmt"

func main() {
    type PhoneNumber string
    type NIP string
    type Age int

    var phoneNumber PhoneNumber = "082222222222"
    var nip NIP = "00000000 000000 0 000"
    var age Age = 22

    fmt.Println(phoneNumber)
    fmt.Println(nip)
    fmt.Println(age)
}