package main

import (
	"fmt"
)

func main() {

	var aray [3]int
	aray[0] = 4
	aray[1] = 5
	fmt.Println(aray)
	a := [3]string{"mujahid1", "mj2", "mj3"}
	fmt.Println(a)
	// dynamic length declaration
	b := [...]string{"mujahid1", "mj2", "mj3", "mj4"}
	c := [...]string{"mujahid1", "mj2", "mj3", "mj4"}
	fmt.Println(b)
	//length of aray

	fmt.Println(len(b))
	//2 arrays comparison
	fmt.Println("a==b", c == b)

	//looping over array
	for index, value := range b {
		fmt.Println(index, value)
	}
}
