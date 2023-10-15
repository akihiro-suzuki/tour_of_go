package main

import (
	"fmt"
	"strconv"
)

func main() {
	var list []string = nil
	fmt.Printf("list: %v\n", list)
	for i := 0; i < 10; i++ {
		list = append(list, strconv.Itoa(i))
	}
	fmt.Println(list)
}
