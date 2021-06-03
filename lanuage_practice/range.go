package main

import "fmt"

func main() {
	//declaring an array of int
	nums := []int{2, 3, 4}
	sums := 0

	for _, num := range nums {
		sums += num
	}

	fmt.Println("Sums: ", sums)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
