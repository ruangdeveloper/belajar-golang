package main

import "fmt"

func main() {
	var name = "Ruang Developer"
	var alias = "Ruang Developer"
	var nikname = "ruangdeveloper"

	var result1 = name == alias
	fmt.Println(result1)

	var result2 = name == nikname
	fmt.Println(result2)

	var result3 = name != nikname
	fmt.Println(result3)
}
