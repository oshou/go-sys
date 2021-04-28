package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	var buffer bytes.Buffer

	reader := bytes.NewBufferString("Example of io.TeeReader\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = ioutil.ReadAll(teeReader)

	fmt.Println(buffer.String())
}
