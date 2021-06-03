package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Second style start here")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Third style start here")
	for i := 0; i < 10; i++ {
		fmt.Println("Loop")
		break
	}
	fmt.Println("Fourth style start here")
	for i := 1; i <= 9; i++ {
		if i%2 == 0 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Ends here")
}
