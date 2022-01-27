package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func samplegofunc(input [2000000]bool) {
	for _, value := range input {
		fmt.Println("go routine1", "Now time is : ", time.Since(start), "printing :", value)
	}

}

func samplegofunc2(input [2000000]bool) {
	for _, value := range input {

		fmt.Println("go routine2", "Now time is : ", time.Since(start), "printing :", value)
	}

}

func main() {

	fmt.Println("Printing main thread")

	var bigaray [2000000]bool

	go samplegofunc(bigaray)
	go samplegofunc2(bigaray)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Ending main thread")

}
