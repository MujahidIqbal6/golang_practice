package main

import "fmt"

func changeValueThroughAdress(ptr *int) {
	*ptr = 99
}

func main() {

	var r int = 100
	fmt.Println(&r)

	var adress = &r
	fmt.Println(adress)

	//string pointer

	var testString = "mujahidAdress"
	var strPointer *string = &testString
	//print adress
	fmt.Println(strPointer)

	//print value
	fmt.Println(*strPointer)

	//change value from a reference
	*strPointer = "new adress for mujahid"

	fmt.Println(testString)

	// new method that returns pointer only

	newTest := new(int)

	fmt.Println(newTest)

	//print value of this newly created
	fmt.Println(*newTest)

	//call function to change value
	changeValueThroughAdress(newTest)

	//print value of this newly created
	fmt.Println(*newTest)
}
