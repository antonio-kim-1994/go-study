package main

import (
	"fmt"
	"time"
)

func main() {
	go sleepyGopher() // goroutine 생성을 위해 go 키워드 사용
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
