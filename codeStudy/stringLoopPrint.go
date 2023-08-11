package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string
	var a string
	fmt.Scan(&s1, &a)
	numA, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result string
	for i := 0; i < numA; i++ {
		result += s1
	}
	fmt.Println(result)
}
