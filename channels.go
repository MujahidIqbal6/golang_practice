package main

import (
	"fmt"
	"sync"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func samplegofunc(input [2000000]bool, chnl chan int) {
	fmt.Println("Now in GoFunction")
	looper := <-chnl
	fmt.Println(looper)

	val, ok := <-chnl

	fmt.Println(val, ok)
}

func squares(chnl chan int) {
	for i := 0; i < 9; i++ {
		chnl <- i * i
	}
	close(chnl)
}

func service1(c chan string) {

	time.Sleep(time.Millisecond)
	c <- "hello from service1"
}

func service2(c chan string) {
	time.Sleep(2 * time.Millisecond)
	c <- "hello from service2"
}

func service3(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "hello from service2"
}

func service4(wg *sync.WaitGroup, value int) {
	time.Sleep(5 * time.Second)
	println("service 4 called on integer value:", value)
	wg.Done()
}

func main() {

	fmt.Println("Printing main thread")
	var bigaray [2000000]bool
	chnl := make(chan int)

	fmt.Println(chnl)
	go samplegofunc(bigaray, chnl)
	chnl <- 100
	fmt.Println("Closing integer channel")
	close(chnl)

	chnl2 := make(chan int)
	go squares(chnl2)

	for val := range chnl2 {
		fmt.Println(val)
	}

	fmt.Println("length of channel is:", len(chnl2))

	fmt.Println("Exiting main thread")

	chnl4 := make(chan string)
	chnl5 := make(chan string)

	go service2(chnl5)
	go service1(chnl4)

	select {
	case res := <-chnl4:
		fmt.Println("response from service 1", res)
	case res := <-chnl5:
		fmt.Println("response from service 2", res)

	case <-time.After(3 * time.Second):
		fmt.Println("times out")
	}

	go service3(chnl4)

	select {
	case res := <-chnl4:
		fmt.Println("response from service 1", res)

	case <-time.After(3 * time.Second):
		fmt.Println("Timed Out because no service has returned any value")
	}

	//wg struct example
	fmt.Println("creating wg instance")
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go service4(&wg, i)
	}

	//wait until all service4 calls have returned
	wg.Wait()
}
