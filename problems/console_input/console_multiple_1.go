package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var inputRange int
	var listVal	list.List
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input the no of value do you want to input: ")
	if _, err := fmt.Scan(&inputRange); err != nil {
		log.Print("Get input failed due to ", err)
		return
	}

	fmt.Println("Input the value: ")

	for i := 1; i <= inputRange; i++ {
		fmt.Println("input value %d: ", i)
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		listVal.PushBack(input)
	}
	for element := listVal.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int64))
	}
}
