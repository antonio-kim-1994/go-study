package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity) // {-1.9462 354.4734}
	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight) // {4.5 135.9}

	// composite literal with values only
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit) // {-14.5684 175.472636}

	// printing keys
	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("%v\n", curiosity)  // {-4.5895 137.4417}
	fmt.Printf("%+v\n", curiosity) // {lat:-4.5895 long:137.4417}

	// copy value
	bradbury := location{-4.5895, 137.4417}
	curiosity1 := bradbury

	curiosity1.long += 0.0106
	fmt.Println(bradbury, curiosity1) // {-4.5895 137.4417} {-4.5895 137.4523}
}
