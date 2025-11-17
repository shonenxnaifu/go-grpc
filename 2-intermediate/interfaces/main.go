package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

type rect1 struct {
	width, height float64
}

func (r rect1) area() float64 {
	return r.height * r.width
}

func (r rect1) perim() float64 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 5, height: 6}

	fmt.Println("---rect---")
	measure(r)
	fmt.Println("---circle---")
	measure(c)
	fmt.Println("---rect 1---")
	measure(r1)
	fmt.Println("circle diameter:", c.diameter())

	myPrinter(1, "john", 45.9, true)
	value(1)
	value("john")
	value(false)
}

func value(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	default:
		fmt.Println("Type: unkown")
	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
