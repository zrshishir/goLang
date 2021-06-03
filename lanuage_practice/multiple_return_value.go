package main

import "fmt"

func vals() (int, int) {
	return 8, 9
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//getting just second parameter
	_, c := vals()
	fmt.Println(c)
}
