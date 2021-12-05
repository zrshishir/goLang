package main

import (
	"fmt"
	"log"
)

func main() {
	var inputNumber int
	var squire, cube int

	fmt.Println("Please input a number: ")

	if _, err := fmt.Scan(&inputNumber); err != nil {
		log.Print("get input is failed due to ", err)
		return
	}

	squire, cube = powerSeriesReturnVar(inputNumber)
	fmt.Println("Squire is: ", squire, ". Cube is: ", cube)
}

func powerSeriesReturnVar(a int) (squire int, cube int) {
	squire = a * a
	cube	= a * a * a
	return
}
