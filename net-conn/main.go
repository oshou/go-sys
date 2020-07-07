package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// Create Connection
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	if _, err = io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"); err != nil {
		log.Fatalln(err)
	}

	// Copy response to os.Stdout
	if _, err = io.Copy(os.Stderr, conn); err != nil {
		log.Fatalln(err)
	}
}
