package main

import (
	"fmt"
	"time"
)

func main() {
	// q: when this tick channel fired?
	// tickは毎回なる(every 100msec)
	tick := time.Tick(100 * time.Millisecond)
	// Afterは１回だけ
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
