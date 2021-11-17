package main

import (
	"fmt"
	"log"
)

func main() {
	var i, j, k int

	if _, err := fmt.Scan(&i, &j, &k); err != nil {
		log.Print("Scan for i, j, k have been failed due to ", err)
		return
	}

	fmt.Println(i, j, k)
}
