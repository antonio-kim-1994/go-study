package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher1()
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher1() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
