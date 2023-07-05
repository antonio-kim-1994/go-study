package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Go의 json 패키지 field의 첫 단어를 대문자로 시작해야 하고, 복합어의 경우 CamelCase로 표기해야 한다.
	// 단, struct tag를 통해 원하는 단어로 표기가 가능하다.
	type location struct {
		Lat, Long float64 // field 명은 반드시 대문자로 시작해야 한다.
	}
	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity) // json.Marshal : Go 값(예를 들어, 구조체, 슬라이스, 맵 등)을 JSON 포맷의 바이트 슬라이스로 변환하는 역할
	exitOnError(err)
	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
