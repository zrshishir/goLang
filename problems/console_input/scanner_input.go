package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("type a number")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Println("you have typed %d ", input)
}
