package main

import (
	"encoding/json"
	"fmt"
	"log"

	erlangc "example.xyz/src/erlangc"
)

// Result godoc
type Result struct {
	Tollbooths, Pressure, Blocking, Queueing int32
}

func main() {
	fmt.Print("Enter the number of tollbooths : ")
	var tollbooths int32
	fmt.Scan(&tollbooths)
	fmt.Println("tollbooths: ", tollbooths)

	fmt.Print("Enter interval in minutes : ")
	var x int32
	fmt.Scan(&x)
	fmt.Println("interval: ", x, "minutes")

	fmt.Print("Enter the number of requests : ")
	var requests int32
	fmt.Scan(&requests)
	fmt.Println("requests: ", requests, " per ", x, " minute(s)")

	fmt.Print("Enter the delay between requests in seconds : ")
	var delay int32
	fmt.Scan(&delay)
	fmt.Println("delay: ", delay)

	// if [requests] occur and it takes [delay] minutes
	// to complete 1 single request, what is the erlang
	// measurement for an [x] second interval?
	result := Result{}
	result.Tollbooths = tollbooths
	result.Pressure = erlangc.Pressure(requests, delay, x)
	result.Blocking = erlangc.B(result.Tollbooths, result.Pressure)
	result.Queueing = erlangc.C(1, result.Pressure)
	prettyPrint(result)
}

func prettyPrint(i Result) {
	prettyJSON, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
