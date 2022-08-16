package main

import (
	"fmt"
	"strings"
)

func formatAndPrint(text string, formatter func(string) string) {
	formatted := formatter(text)
	fmt.Println(formatted)
}

func upper(text string) string {
	return strings.ToUpper(text)
}

func lower(text string) string {
	return strings.ToLower(text)
}

func main() {
	formatAndPrint("hello world", upper)
	formatAndPrint("HELLO WORLD", lower)
}
