package main

import (
	"bufio"
	"os"
)

func main() {
	buf := bufio.NewWriter(os.Stdout)
	if _, err := buf.WriteString("bufio.Writer\n"); err != nil {
		panic(err)
	}

	if err := buf.Flush(); err != nil {
		panic(err)
	}

	if _, err := buf.WriteString("example\n"); err != nil {
		panic(err)
	}

	if err := buf.Flush(); err != nil {
		panic(err)
	}
}
