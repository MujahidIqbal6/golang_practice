package main

import (
	"fmt"
)

var globalVariable = 500

func init() {
	fmt.Println("initializer function of package called")
	globalVariable = 600
}

func main() {

	//creating a new type float from float64
	type float float64

	//if else comparison test
	var f float = 52.2
	var g float64 = 52.2
	if f == 52.2 {
		fmt.Println(g)
	} else {
		fmt.Println("its not g")
	}

	//for loop test
	fmt.Println("for loop test")
	for i := 8; i < 16; i++ {
		fmt.Println(i)
	}

	j := 2
	for j < 17 {
		fmt.Println(j)
		j++
	}

	fmt.Println("for loop test with another approach")
	for j > 16 {
		fmt.Println(j)
		j++
		if j == 100 {
			break
		}
	}

	//print global variable defined in package
	fmt.Println("Printing globally defined variable in Package")
	fmt.Println(globalVariable)

	//testing strings
	var testString string

	testString = "testing"
	fmt.Println("string is:" + testString)
	fmt.Print("length of string is : ")
	fmt.Println(len(testString))

	//looping through test string to check what it prints
	for i := 0; i < len(testString); i++ {
		fmt.Printf("%c ", testString[i])
	}
}
