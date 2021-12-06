package main

import (
	"fmt"
	"log"
)

func main() {
	var inputNumber int
	var squire int
	var cube int

	fmt.Println("Please input for squire and cube")

	if _, err := fmt.Scan(&inputNumber); err != nil {
		log.Print("Get input if failed due to ", err)
		return
	}

	squire, cube = powerSeries(inputNumber)

	fmt.Println("Squire is: ", squire, " And cube is : ", cube)
}

func powerSeries(input int) (int, int) {
	return input * input, input * input * input
}
