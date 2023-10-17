package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(bytes []byte) (int, error) {
	readBytes, err := r.r.Read(bytes)
	for i := 0; i < len(bytes); i++ {
		if 65 <= bytes[i] && bytes[i] <= 90 {
			bytes[i] = bytes[i] + 13
			if bytes[i] > 90 {
				bytes[i] = bytes[i] - 90 + 65 - 1
			}
			continue
		}
		if 97 <= bytes[i] && bytes[i] <= 122 {
			bytes[i] = bytes[i] + 13
			if bytes[i] > 122 {
				bytes[i] = bytes[i] - 122 + 97 - 1
			}
			continue
		}
	}
	return readBytes, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
