package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buf.String())
}
