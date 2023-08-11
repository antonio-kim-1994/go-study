package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type FloatTest struct {
	TotalMemory, UsedMemory, AvailableMemory, MemoryUsage float64
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	test := FloatTest{
		TotalMemory:     12.884901888,
		UsedMemory:      3.270585976,
		AvailableMemory: 9.614315912,
		MemoryUsage:     25.38308793057998,
	}

	a := fmt.Sprintf("*전체 메모리량*:\n%.2f GB", test.TotalMemory)
	b := fmt.Sprintf("*메모리 사용량*:\n%.2f GB", test.UsedMemory)
	c := fmt.Sprintf("*사용 가능 메모리량*:\n%.2f GB", test.AvailableMemory)
	d := fmt.Sprintf("*메모리 사용률*:\n%.2f %%", test.MemoryUsage)

	log.Printf("%v\n%v\n%v\n%v", a, b, c, d)
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04")

	fmt.Printf("[%v] Total Memory : %.2f GB\n", formattedTime, test.TotalMemory)
	fmt.Printf("[%v] User Memory : %.2f GB\n", formattedTime, test.UsedMemory)
	fmt.Printf("[%v] Available Memory : %.2f GB\n", formattedTime, test.AvailableMemory)
	fmt.Printf("[%v] Memory Usage : %.2f %%\n", formattedTime, test.MemoryUsage)
}
