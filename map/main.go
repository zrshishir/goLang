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
	//to delete a key
	delete(mp, "pear")
	fmt.Println(mp)
	//importnt tricks: if a key exits ok will be true and val will be assigned with key value
	//if not ok will be false and val will be 0
	val, ok := mp["apple"]
	fmt.Println(val, ok)
}
