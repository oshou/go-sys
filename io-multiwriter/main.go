package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	if _, err := io.WriteString(writer, "io.MultiWriter example\n"); err != nil {
		panic(err)
	}
}
