package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringer is an interface type that describes objects that can describe themselves as strings.
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"ABC", 42}
	b := Person{"DEF", 9001}
	fmt.Println(a, b)
}
