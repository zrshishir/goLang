package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) peri() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.peri())

	rp := &r
	fmt.Println("PArea: ", rp.area())
	fmt.Println("PPerimeter: ", rp.peri())
}
