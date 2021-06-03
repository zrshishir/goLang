package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("taking a value from map: ", v1)

	fmt.Println("length of the map: ", len(m))

	//deleting a key from a map

	delete(m, "k2")

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

}
