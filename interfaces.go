package main

import (
	"fmt"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type naming interface {
	showName() string
}

type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rect) showName() string {
	return "this is a rectangle"
}

func (r rect) Area() float64 {
	return r.height * r.width
}

func (r rect) Perimeter() float64 {
	return r.height + r.width
}

func (c circle) Area() float64 {
	return 6.16 * c.radius
}

func (c circle) Perimeter() float64 {
	return 6.16 + c.radius
}

func tellType(i interface{}) {
	switch i.(type) {

	case string:
		fmt.Println("given type is string")

	case int:
		fmt.Println("given type is integer")

	default:
		fmt.Println("unknown type")
	}

}

func main() {

	var s shape
	s = rect{3.14, 2.3}

	fmt.Println(s)
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

	s = circle{50}
	fmt.Println(s)
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())

	var nameobject naming
	nameobject = rect{4, 4}
	fmt.Println(nameobject.showName())

	//check type switching

	tellType("asd")
	tellType(23)
	tellType(s)

}
