package main

import "fmt"

type Employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

type Company struct {
	string
	int
	bool
	Employee
}

//interface function for struct

func (e Employee) multiplySalary(factor int) int {
	return e.salary * factor
}

//anoynamous struct

func main() {

	var mujahid Employee
	//we will print empty object with zero values
	fmt.Println(mujahid)

	//now set some properties of objects

	mujahid.firstName = "mujahid"
	mujahid.lastName = "iqbal"
	mujahid.salary = 50000
	mujahid.fullTime = false

	//we will print again
	fmt.Println(mujahid)

	ross := Employee{
		firstName: "ross", lastName: "robert", salary: 400, fullTime: true,
	}

	//we will print again
	fmt.Println(ross)

	ross2 := Employee{"amir", "sohail", 3, false}

	//we will print again
	fmt.Println(ross2)

	//pointer to a struct
	ross3 := &Employee{"ross", "sohailrobert", 3, false}

	//we will print again
	fmt.Println(*&ross3.lastName)

	com := Company{"netsol", 4, false, ross2}
	fmt.Println(com)

	//accessing nested structs

	fmt.Println(com.Employee.firstName + " works in " + com.string)

	//call interface func
	fmt.Println(com.Employee.multiplySalary(5))

	//comparing 2 structs, will return false because values are different
	fmt.Println(ross == ross2)

}
