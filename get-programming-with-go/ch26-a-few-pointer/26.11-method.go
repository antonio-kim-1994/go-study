package main

import "fmt"

type personMethod struct {
	name string
	age  int
}

func main() {
	terry := &personMethod{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v\n", terry) // &{name:Terry age:16}

	nathan := personMethod{
		"Nathan",
		17,
	}
	nathan.birthday()
	// 메서드에 포인터 수신자를 사용하면, 그 메서드는 포인터 타입 또는 값 타입 모두에 대해 호출될 수 있다.
	// Go는 자동으로 필요한 경우 값 타입을 포인터 타입으로, 또는 그 반대로 변환한다.
	// birthday 메서드에서 포인터를 요구하고 있어 자동으로 값을 포인터로 변환하여 메서드를 호출한다.
	fmt.Printf("%+v\n", nathan) // {name:Nathan age:18}
}

func (p *personMethod) birthday() {
	p.age++
}
