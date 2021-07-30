package main

import "fmt"

func main(){

	//one way to declare the map
	var mp map[string]int = map[string]int{
		"apple": 5,
		"pear": 6,
		"guava": 9,
	}
	//other way to declare the map
	//mp := make(map[string]int)
	fmt.Println(mp["apple"])
	mp["pear"] = 600
	fmt.Println(mp["pear"])
}
