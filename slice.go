package main

import "fmt"

func main() {

	var fruits = [...]string{
		"Apel",
		"Nanas",
		"Melon",
		"Semangka",
		"Jeruk",
		"Pepaya",
		"Anggur",
	}

	var slice = fruits[2:5]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	chars := [...]string{"A", "B", "C", "D", "E", "F"}

	charSlice1 := chars[4:]
	fmt.Println(chars)      // [A B C D E F]
	fmt.Println(charSlice1) //[E F]

	charSlice2 := append(charSlice1, "G")
	charSlice2[0] = "E lagi"
	fmt.Println(charSlice2) // [E lagi F G]
	fmt.Println(chars)      // [A B C D E F]

	days := make([]string, 2, 6)
	days[0] = "Senin"
	days[1] = "Selasa"

	fmt.Println(days)
	fmt.Println(len(days))
	fmt.Println(cap(days))

	copyDays := make([]string, len(days), cap(days))
	copy(copyDays, days)
	fmt.Println(copyDays)
}
