package main

import "fmt"

func squareAll(array [5]int) {

	for index, value := range array {
		array[index] = value * value
	}

	fmt.Println(array)
}

func main() {

	var slice []int
	//empty slice
	fmt.Println(slice)
	//random array
	testAray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(testAray)

	//creating slice from aray because slice is just a referrence from aray

	slice = testAray[2:5]

	//print slice now

	fmt.Println(slice)

	//change values in original aray now

	testAray[2] = 22

	fmt.Println(slice)

	//length of slice

	fmt.Println(len(slice))

	//change slice again

	slice = testAray[1:2]
	fmt.Println(len(slice))

	//change slice, this change will reflect in aray too

	slice[0] = 33

	fmt.Println(testAray)

	//append func

	newS := append(slice, 44, 55, 66)

	fmt.Println(newS)
	//print original aray
	fmt.Println(testAray)

	newS = append(slice, 101, 102, 103, 104)

	//above step has extended aray beyond its size
	fmt.Println(newS)

	squareAll(testAray)

	//printing arry again to check values of testarray
	fmt.Println(testAray)
}
