package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("format-file-writer.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "string is %s, integer is %d, float is %f", "aaa", 12, 3.5)
}
