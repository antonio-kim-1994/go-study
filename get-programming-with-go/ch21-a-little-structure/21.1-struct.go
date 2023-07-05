package main

import "fmt"

func main() {
	// 여러 필드를 묶어서 새로운 타입을 정의하는 복합 타입이다.
	// 구조체를 사용하면 서로 다른 타입의 데이터를 하나의 논리적 그룹으로 묶을 수 있다.
	var curiosity struct {
		lat  float64
		long float64
	}
	curiosity.lat = -4.5895
	curiosity.long = 137.4417
	fmt.Println(curiosity.lat, curiosity.long) // -4.5895 137.4417
	fmt.Println(curiosity)                     // {-4.5895 137.4417}
}
