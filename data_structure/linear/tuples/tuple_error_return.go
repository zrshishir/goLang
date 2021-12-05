package main

import (
	"fmt"
	"log"
)

func main() {
	var inputNumber int
	var squire, cube int

	fmt.Println("Please input a value: ")

	if _, err := fmt.Scan(&inputNumber); err != nil {
		log.Print("get input failed due to ", err)
		return
	}

	squire, cube, _ = powerSeriesErrReturn(inputNumber)

	fmt.Println("Squire is: ", squire, ". Cube is: ", cube)
}

func powerSeriesErrReturn(a int) (int, int, error) {
	return a * a, a * a * a, nil
}
