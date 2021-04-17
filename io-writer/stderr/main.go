package main

import "os"

func main() {
	os.Stderr.Write([]byte("os.Stderr example\n"))
}
