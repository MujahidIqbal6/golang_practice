package main

import "fmt"

type Organ struct {
	name string
}
type Human struct {
	firstName, lastName string
	organ               Organ
}

type Employee struct {
	firstName, lastName string
}

//new derived type

type Mystring string

//method for employee

func (e Employee) fullName() string {
	return e.firstName + " " + e.lastName
}

//same method name for struct human
func (h Human) fullName() string {
	return h.firstName + " " + h.lastName + " is a human"
}

func (h *Human) changeName() {
	(*h).firstName = "MR. " + h.firstName
}

//promoted method
func (h *Organ) changeOrgan(new string) {
	(*h).name = new
}

func (s *Mystring) addexclaimation() {
	*s = *s + "!"
}

//anoynamous struct

func main() {

	e := Employee{"Mujahid", "Iqbal"}
	fmt.Println(e.fullName())
	h := Human{"amir", "sohail", Organ{"heart"}}
	fmt.Println(h.fullName())
	//changing name of original obj
	h.changeName()
	fmt.Println(h.fullName())
	//print organ
	fmt.Println(h.organ.name)
	//new organ
	h.organ.changeOrgan("stomach")
	//print organ
	fmt.Println(h.organ.name)
	var str Mystring = "hi"
	str.addexclaimation()

	fmt.Println(str)

}
