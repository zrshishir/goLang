package main

import "fmt"

func val(ival int) {
	ival = 3
}

func pointer(iptr *int) {
	*iptr = 4
}

func main() {
	i := 1
	fmt.Println("initial value: ", i)
	val(i)
	fmt.Println("calling the val: ", i)
	pointer(&i)
	fmt.Println("calling pointer: ", i)
	fmt.Println("calling pointer: ", &i)
}
