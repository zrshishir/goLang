package main

import (
	"fmt"
	"log"
)

func main()  {
	var i int

	if _, err := fmt.Scan(&i); err != nil {
		log.Print(" Scan for i has been failed due to ", err)
		return
	}

	fmt.Println(i)
}
