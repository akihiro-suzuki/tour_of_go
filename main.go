package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// q: why this is pointer receiver?
// a: because we want to modify the value of MyError
// q: but below code don't modify the value of MyError
// a: yes, but it's a good practice to use pointer receiver
// q: why?
// a: because it's more efficient
// q: how?
// a: because it's more efficient to pass a pointer than a value
// q: but it leads to confusion. because we don't modify the value of MyError
// a: yes, but it's a good practice to use pointer receiver
// q: why?
// a: because it's more efficient
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// q: ポインタで返さないと、goのerrorインターフェースを満たさない?
// a: そう。errorインターフェースは、ポインタで返すように定義されている
// q: そしてgoでは、ポインタレシーバが推奨されている
// a: そう。ポインタレシーバの方が、値レシーバよりも効率が良いから
// q: つまり、interfaceを実装した型を戻り値にしたいときは、基本的にポインタで返すべき?
// a: そう。ポインタで返すべき

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
