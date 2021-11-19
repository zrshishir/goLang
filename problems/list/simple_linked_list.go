package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

// @ Todo: https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem?isFullScreen=true

func main()  {
	var iterator, i int64
	var listVal list.List

	fmt.Println("Input iterator")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	iterator, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	for i = 1; i <= iterator; i++ {
		scanner.Scan()
		input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		listVal.PushBack(input)
	}

	fmt.Println("Output")
	for element := listVal.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int64))
	}
}