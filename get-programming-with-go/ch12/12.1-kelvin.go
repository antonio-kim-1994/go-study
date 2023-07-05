package main

import "fmt"

func KelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := KelvinToCelsius(kelvin)
	fmt.Print(kelvin, "° K is ", celsius, "° C")
}
