package main

import (
	"fmt"
)

func divide(a, b int) int {

	defer deferFromdivide()
	if b == 0 {
		panic("Cannot divide by zero(0)")
		return 0
	} else {
		return (a / b)
	}

}

func deferFromdivide() {

	fmt.Println("this should be called second  from divide func stack")

	r := recover()

	fmt.Println("recover func value", r)
}

func deferFromMain() {

	fmt.Println("this should be called first from main stack")
}

func main() {

	defer deferFromMain()
	fmt.Println(divide(27, 0))
	fmt.Println("main ended")
}
