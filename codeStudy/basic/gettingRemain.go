package main

import "fmt"

func gettingRemain(num1 int, num2 int) int {
	return num1 % num2
}

func main() {
	fmt.Println(gettingRemain(1, 2))
}
