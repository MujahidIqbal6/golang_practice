package main

import (
	"fmt"
)

func main() {

	//it will show every argument by inserting a space in between
	fmt.Print(1, 2, 3, 4, 5, false, true, 2, 33, 'e')
	fmt.Print("\n Hey!", "how are You", "This all words will be concatenaded", "\n")

	//println will alwats contatenate
	fmt.Println("hey!", "how are you?", false, 44)
	fmt.Println("hey!", "I am fine!", true, 22.22)

	//sprint is same as print, only it returns

	name := fmt.Sprint("hi this is mujahid!")
	fmt.Println(name)

	//printing with parameters, %v is a parameter

	fmt.Printf("hi my name is %v \n", "mujahid iqbal")

	//creating dynamic string

	new_sentence := fmt.Sprintf("hi my name is %v and i am %v years old and i am currently %v", "mujahid", "26", "single")
	fmt.Println(new_sentence)

	//printing struct

	anoynamous_struct := struct {
		name string
		age  int
	}{"mujahid", 26}

	//simple print will just print with values
	fmt.Println(anoynamous_struct)

	//using printf + format will print struct with vaiable names too
	fmt.Printf("%+v \n", anoynamous_struct)

	// %T is used to show type

	fmt.Printf("%T \n", anoynamous_struct)
	fmt.Printf("%T \n", new_sentence)

	dd := fmt.Sprintf("%T \n", anoynamous_struct)
	fmt.Println(dd)
}
