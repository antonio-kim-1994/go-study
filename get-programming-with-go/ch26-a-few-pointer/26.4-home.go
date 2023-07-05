package main

import "fmt"

func main() {
	canada := "Canada"
	var home *string                   // 문자열의 포인터를 가리킨다. 초기화되지 않은 상태이므로 nil 값을 갖는다.
	fmt.Printf("home is a %T\n", home) // home is a *string

	home = &canada
	fmt.Println(*home) // Canada
}
