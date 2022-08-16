package main

import "fmt"

func main() {
	website := make(map[string]string)

	website["name"] = "Ruang Developer"
	website["domain"] = "www.ruangdeveloper.com"
	fmt.Println(website)

	// Menghapus data domain
	delete(website, "domain")
	fmt.Println(website)
}
