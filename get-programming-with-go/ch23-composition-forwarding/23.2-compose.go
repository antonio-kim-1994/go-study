package main

import "fmt"

/*
타입 선언을 사용하는 이유
가독성: 각각의 타입과 필드명이 가지는 의미를 명확하게 나타낸다. 예를 들어, location 구조체는 위도와 경도를 표현하며, temperature는 고온과 저온을 표현
안정성: 각 타입에 대한 연산이나 메소드를 정의함으로써 코드의 안정성을 향상시킨다. 예를 들어, celsius와 같은 사용자 정의 타입에는 특정한 메소드를 연결할 수 있으며, 이를 통해 해당 타입의 값이 어떻게 조작되어야 하는지를 명확하게 정의할 수 있다.
캡슐화와 조직화: 데이터와 그에 관련된 로직을 함께 묶어서 조직화하고 캡슐화한다. report와 같은 복합 타입은 sol, temperature, location 등의 여러 값을 하나로 묶어 관리하므로 관련 데이터와 로직을 쉽게 조직화할 수 있다.
재사용성: temperature, location, celsius와 같은 타입을 재사용하면서 코드 중복을 줄일 수 있다. 이를 통해 코드의 일관성을 유지하고, 변경이 필요할 때 수정이 용이하다.
*/
type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)                            // {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5895 long:137.4417}}
	fmt.Printf("a balmy %v° C\n", report.temperature.high) // a balmy -1° C
	fmt.Printf("average %v° C\n", t.average())             // average -39.5° C
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
