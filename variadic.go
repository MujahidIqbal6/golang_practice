package main

import "fmt"

func getMultiples(factor int, args ...int) []int {
	multiples := make([]int, len(args))
	for index, val := range args {
		multiples[index] = val * factor
	}
	return multiples
}

func getMultiplesinPlace(factor int, args ...int) []int {

	for index, val := range args {
		args[index] = val * factor
	}
	return args
}

func morethanone(factor int, divisor int, args ...int) int {

	sum := 0
	for _, val := range args {
		sum = sum + val
	}
	return sum
}

func main() {

	s := []int{10, 20, 30}
	mult1 := getMultiples(2, s...)
	mult2 := getMultiples(3, 1, 2, 3, 4)
	fmt.Println(mult1)
	fmt.Println(mult2)

	s1 := []int{10, 20, 30}
	mult3 := getMultiplesinPlace(2, s1...)
	fmt.Println(mult3)
	fmt.Println(morethanone(1, 1, 5, 5, 5, 5))

}
