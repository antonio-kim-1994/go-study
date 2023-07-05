package main

import (
	"fmt"
	"time"
)

// channel : 고루틴(Goroutine) 간에 데이터를 전송하고 동기화하는 수단
// 고루틴 간에 데이터를 공유하는 대신 채널(Channel)을 통해 데이터를 주고받는다.
// 즉, 공유 메모리를 통해 통신하는 대신, 통신을 통해 메모리를 공유한다.
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher3(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher3(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id // send a value back to main
}
