package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"
	fmt.Println(planets)       // map[Earth:whoops Mars:Sector ZZ9]
	fmt.Println(planetsMarkII) // map[Earth:whoops Mars:Sector ZZ9]
	delete(planets, "Earth")
	fmt.Println(planetsMarkII) // map[Mars:Sector ZZ9]

	/*
		Map은 Array와 다르게 새로운 배열을 생성하여 값을 복사하지 않는다.
		동일한 값을 공유하기 때문에 한 쪽의 값이 변경될 경우 참조한 변수의 값도 바뀌게 된다.
	*/
}
