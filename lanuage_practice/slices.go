package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("out: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("length: ", len(s))

	s = append(s, "d")
	s = append(s, "m", "k")
	fmt.Println("append: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	l := s[2:5]
	fmt.Println("from:to", l)

	l = s[:5]
	fmt.Println("till: ", l)

	l = s[2:]
	fmt.Println("from: ", l)
}
