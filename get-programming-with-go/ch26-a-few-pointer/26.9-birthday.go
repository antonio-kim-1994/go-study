package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca) // {name:Rebecca superpower:imagination age:15}
}

// 기본적으로 golang에서 함수나 메서드를 통해 argument를 받을 때 argument의 값을 복사해서 사용한다.
func birthday(p *person) { // argument로 메모리 주소값을 받고, 해당 메모리 주소가 가리키는 값을 변경한다.
	p.age++
}
