package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Johm", age: 29})
	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("JohnD"))

	s := person{name: "Sean", age: 30}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.age)
	sp.age = 35
	fmt.Println(sp.age)
}
