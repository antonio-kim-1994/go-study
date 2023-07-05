package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4:4] // length 4, capacity 4
	worlds := append(terrestrial, "Ceres")
	fmt.Println(worlds)
	fmt.Println(planets)
	terrestrial = planets[0:4] // length 4, capacity 8
	worlds = append(terrestrial, "Ceres")
	fmt.Println(worlds)
	fmt.Println(planets)
}
