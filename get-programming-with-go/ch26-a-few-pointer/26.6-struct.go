package main

import "fmt"

func main() {
	type person struct {
		name, superpower string
		age              int
	}
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	timmy.superpower = "flying"
	// (*timmy).superpow = "flying" 과 동일
	fmt.Printf("%+v\n", timmy) // &{name:Timothy superpower:flying age:10}
}
