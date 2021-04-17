package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}

	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"

	if _, err := io.WriteString(writer, "gzip.Writer example\n"); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}
}
