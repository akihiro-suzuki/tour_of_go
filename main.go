package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

// q: * つけるとmutableになるという理解でいいのかな
// a: そうですね。ポインタを使うと、呼び出し元の値を更新できます。
// q: & をつける例はありますか？
// a: あまりないですね。ポインタを使うときは、呼び出し元の値を更新したいときに使います。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
