package main

import (
	"bufio"
	"os"
)

var bufsize = 64 * 1024

func main() {
	var y []byte
	y = []byte("y\n")
	yLen := len(y)
	buf := make([]byte, bufsize)
	for i := 0; i < len(buf)-yLen; i += yLen {
		for s := 0; s < len(y); s++ {
			buf[i+s] = y[s]
		}
	}

	f := bufio.NewWriterSize(os.Stdout, bufsize)
	defer f.Flush()
	for {
		f.Write(buf)
	}
}