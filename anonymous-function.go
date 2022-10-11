package main

import "fmt"

func authorized(username string, authorize func(string) bool) {
	if authorize(username) {
		fmt.Println("Authorized")
	} else {
		fmt.Println("Unauthorized")
	}
}

func main() {
	// Contoh 1 - Dengan variabel
	onlyAdmin := func(username string) bool {
		return username == "admin"
	}
	authorized("admin", onlyAdmin)

	// Contoh 2 - Langsung di parameter
	authorized("admin", func(username string) bool {
		return username == "member"
	})
}
