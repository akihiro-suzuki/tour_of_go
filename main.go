package main

import "fmt"

// q: X, Yのデフォルト値を設定しておきたい
type MyStruct struct {
	X, Y int
}

func main() {
	var i interface{} = "hello"

	m, ok := i.(MyStruct) // {0, 0}, false
	fmt.Println(m, ok)

	m2 := i.(MyStruct) // panic
	fmt.Println(m2)
}
