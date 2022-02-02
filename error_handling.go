package main

import (
	"errors"
	"fmt"
)

type PathError struct {
	path string
}

func (e *PathError) Error() string {

	return fmt.Sprintf("Error occurred while reading path %v", e.path)
}

func throwError() error {

	return &PathError{path: "c:\\dektop\\mujahid"}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return (a / b), nil
	}

}

func main() {

	fmt.Println("main started")
	if result, Error := divide(20, 0); Error == nil {
		println("answer is :", result)
	} else {

		fmt.Println("Error: ", Error)
	}

	//creating a custom error type

	//getting type of error
	err := throwError()
	fmt.Println(err)
	switch e := err.(type) {
	case *PathError:
		fmt.Println("error is of type path")
	default:
		fmt.Println(e)
	}
}
