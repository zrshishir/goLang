package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}

func main() {
	fmt.Println("5 + 4 = ", plus(5, 4))
	fmt.Println("5 - 4 = ", minus(5, 4))
}
