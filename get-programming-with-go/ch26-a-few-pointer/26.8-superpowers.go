package main

import "fmt"

func main() {
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	// Go가 배열 및 슬라이스에 대한 포인터를 자동으로 역참조한다.
	// 변수의 값을 변경하려는 경우 (예: *p = 10와 같이)에는 여전히 역참조가 필요하다.
	fmt.Println(superpowers[0])   // flight
	fmt.Println(superpowers[1:2]) // [invisibility]
}
