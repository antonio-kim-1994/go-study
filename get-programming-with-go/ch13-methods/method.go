package main

import "fmt"

//type celsius float64
//type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var (
		k kelvin = 294.0
		c celsius
	)

	c = k.celsius()
	fmt.Println(c)
}
