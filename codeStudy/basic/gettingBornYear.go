package main

import "fmt"

func solution(age int) int {
	bornYear := 2022 - age + 1
	return bornYear
}

func main() {
	fmt.Println(solution(23))
}
