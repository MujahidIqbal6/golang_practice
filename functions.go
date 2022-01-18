package main

import (
	"fmt"
)

func sum(a int, b int) {
	fmt.Println(a + b)
}

func subtract(a int, b int) int {
	return a - b
}

func sumall(a, b, c, d int) {
	fmt.Println(a + b + c + d)
}

func returnTest(firstName, lastName string) string {
	fullname := fmt.Sprintf("%s%s", firstName, lastName)

	return fullname
}

func returnMultipleTest(first, second int) (int, int) {

	a := first - 1
	b := second + 1

	return a, b
}

func deferfunc() {
	fmt.Println("this is a sample deferred function")
}

func funcwithParam(a int, b int, f func(int, int) int) int {

	fmt.Println("function passed as parameter called")
	c := f(a, b)
	return c
}

func main() {

	sum(5, 6)
	defer deferfunc()
	sumall(1, 1, 1, 1)
	fmt.Println(returnTest("mujahid ", "iqbal"))
	a, b := returnMultipleTest(101, 99)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(funcwithParam(5, 10, subtract))
}
