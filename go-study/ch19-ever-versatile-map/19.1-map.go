package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v° C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon) // 0

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	/*
		comma idiom 패턴 : map에서 key를 이용해 value를 가져올 때 두 개의 값을 반환하는 것을 활용
		첫 번째 값은 키에 대응하는 값이며, 두 번째 값은 키가 맵에 존재하는지를 나타내는 불리언 값
		만약 "Moon"이라는 키가 temperature 맵에 있다면, ok는 true가 되고, moon은 그에 대응하는 값을 갖는다.
		"Moon"이라는 키가 temperature 맵에 없다면, ok는 false
		if moon, ok := temperature["Moon"]에서 ok는 새로 생성하는 변수 명이기 때문에 다른 변수명을 사용해도 무방하다.
	*/
}
