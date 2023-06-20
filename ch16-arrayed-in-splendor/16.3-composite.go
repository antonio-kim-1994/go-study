package main

import "fmt"

func main() {
	planets := [...]string{ // Go compiler counts the elements
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	for i, planet := range planets {
		fmt.Println(i, planet)
	}
}
