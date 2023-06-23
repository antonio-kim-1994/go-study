package main

import (
	"fmt"
	"sort"
)

func main() {
	// Go는 기본적으로 Set 타입을 지원하지 않는다.
	// 하지만, map 타입을 통해 Set의 기능을 구현할 수 있다.
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}
	if set[-28.0] {
		fmt.Println("set member")
	}
	fmt.Println(set)

	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique) // [-33 -31 -29 -28 -23 32]
}
