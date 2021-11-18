package main

import (
	"container/list"
	"fmt"
	"log"
)

func main() {
	var inputRange int
	var listVal	list.List
	var intVal int

	fmt.Println("Please input the no of value do you want to input: ")
	if _, err := fmt.Scan(&inputRange); err != nil {
		log.Print("Get input failed due to ", err)
		return
	}


	fmt.Println("Input the value: ")
	for i := 1; i <= inputRange; i++ {
		fmt.Println("input value: ")
		if _, err := fmt.Scan(&intVal); err != nil {
			log.Print("failed to input due to ", err)
			return
		}
		listVal.PushBack(intVal)
	}
	fmt.Println(listVal)
}
