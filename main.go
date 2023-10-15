package main

import "fmt"

func do(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", t, t*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", t, len(t))
	default:
		fmt.Printf("I don't know about type %T!\n", t)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	var i interface{} = nil
	do(i)
}
