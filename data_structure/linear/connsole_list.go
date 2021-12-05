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
	var listStack list.List
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input how many value do you want to input: ")
	if _, err := fmt.Scan(&inputRange); err != nil {
		log.Print("get input failed due to ", err)
		return
	}

	for i := 1; i <= inputRange; i++ {
		fmt.Println("Input value %d", i)
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		listStack.PushBack(input)
	}

	fmt.Println("Output are: ")
	for element := listStack.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int64))
	}
}
