package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	}

	fmt.Println("In case of time using switch")

	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is not weekend")
	}

	fmt.Println("Another style of switch statement")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	fmt.Println("Another style of using switch statement")

	WhoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am an integer")
		default:
			fmt.Println("I don't know my type ", t)
		}
	}

	WhoAmI(true)
	WhoAmI(2)
	WhoAmI("Shishir")

}
