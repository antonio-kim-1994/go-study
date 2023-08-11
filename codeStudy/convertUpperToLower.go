package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "aBcDeFg"

	var result string
	for _, char := range input {
		if unicode.IsUpper(char) {
			result += string(unicode.ToLower(char))
		} else {
			result += string(unicode.ToUpper(char))
		}
	}

	fmt.Println(result)
}
