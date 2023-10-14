package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())

	f := MyFloat(3.0)
	fmt.Println(f.Abs())

	f2 MyFloat := -3.0
	fmt.Println(f2)

}
