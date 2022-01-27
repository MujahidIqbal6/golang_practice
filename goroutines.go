package main

import (
	"fmt"
	"time"
)

func samplegofunc() {
	fmt.Println("This is a sample Go Func")
}

func samplego2func() {
	fmt.Println("This is a sample Go Func for phase 2")
	time.Sleep(time.Millisecond)
}

func main() {

	fmt.Println("Printing main thread")

	go samplegofunc()
	go samplego2func()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Ending main thread")

}
