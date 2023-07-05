package main

import "fmt"

func main() {
	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) // Christopher J. Scolese

	bolden := "Charles F. Bolden"
	administrator = &bolden     // &scolese -> &bolden으로 메모리 주소 변경
	fmt.Println(*administrator) // Charles F. Bolden

	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator) // Charles Frank Bolden Jr.

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden) // Maj. Gen. Charles Frank Bolden Jr.

	major := administrator                            // -> &bolden
	*major = "Major General Charles Frank Bolden Jr." // -> &bolden
	fmt.Println(bolden)                               // Major General Charles Frank Bolden Jr.
	fmt.Println(*administrator)                       // Major General Charles Frank Bolden Jr.
	fmt.Println(administrator == major)               // true

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot          // &bolden -> &lightfoot으로 메모리 주소 변경
	fmt.Println(administrator == major) // false

	charles := *major         // string 값을 할당함으로써 값 복사 발생
	*major = "Charles Bolden" // -> &bolden
	fmt.Println(charles)      // Major General Charles Frank Bolden Jr.
	fmt.Println(bolden)       // Charles Bolden
}
