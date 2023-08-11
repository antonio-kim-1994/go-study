package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ""
	parsedData := strings.Split(input, "&")

	result := make(map[string]string)

	for i := range parsedData {
		data := strings.Split(parsedData[i], "=")
		result[data[0]] = data[1]
		i++
	}

	fmt.Println(result)
}
