package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	input := "appadmin-web-build-dev"
	segments := strings.Split(input, "-")

	result := make(map[string]string)
	result["serviceName"] = segments[0]
	result["serviceType"] = segments[1]
	result["command"] = segments[2]

	if len(segments) == 4 {
		result["updateEnv"] = segments[3]
	} else if len(segments) == 5 {
		result["updateEnv"] = segments[3]
		result["updateType"] = segments[4]
	}

	if result["updateType"] == "" {
		fmt.Println("Length : 4")
		fmt.Printf("%v\n", time.Now().Format("060102-1504"))
	} else {
		fmt.Println("Length : 5")
	}

	fmt.Println("End")
}
