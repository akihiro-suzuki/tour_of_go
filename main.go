package main

import (
	"fmt"
	"math"
)

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = &f // a *MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	a = &v

	fmt.Println(a.Abs())
	describe(a)
}

func describe(a Abser) {
	// %T はtype
	// %v はvalue
	writtenBytes, err := fmt.Printf("(%v, %T)\n", a, a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(writtenBytes)
	}
}
