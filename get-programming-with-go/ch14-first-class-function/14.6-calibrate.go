package main

import "fmt"

type kelvin3 float64
type sensor3 func() kelvin3

func realSensor3() kelvin3 {
	return 0
}

func calibrate(s sensor3, offset kelvin3) sensor3 {
	return func() kelvin3 {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor3, 5)
	fmt.Println(sensor())
}
