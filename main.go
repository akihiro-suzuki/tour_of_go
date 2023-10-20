package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// コメントアウトすると、rangeでdeadlockになる
	close(c)
}

func main() {
	c := make(chan int, 5)
	fmt.Println(cap(c))
	go fibonacci(cap(c), c)

	// closeすると、rangeが終了する
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("done")
}
