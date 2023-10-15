package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (ns ErrNegativeSqrt) Error() string {
	f := float64(ns)
	return fmt.Sprintf("cannot Sqrt negative number: %v", f)
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	} else {
		// q: why this is invalid operation?
		// a: because ErrNegativeSqrt is not a pointer
		return 0, ErrNegativeSqrt(x)
	}
}
func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
