package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Clear is better than clever.")
	buffer := make([]byte, 8)
	for {
		n, err := reader.Read(buffer)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, buffer)
		// %q は Go の構文で安全にエスケープされた単一引用符で囲まれた文字リテラルを意味します
		fmt.Printf("buffer[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			break
		}
	}
}
