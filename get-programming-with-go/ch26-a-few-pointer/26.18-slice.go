package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"pluto",
	}
	reclassify(&planets)
	fmt.Println(planets) // [Mercury Venus Earth Mars Jupiter Saturn Uranus Neptune]
}

// parameter로 slice의 포인터를 받아 내부 배열을 첫 8개 요소만 저장하도록 수정한다.
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}
