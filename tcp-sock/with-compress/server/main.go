package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// start server
	server, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	for {
		// accept connection
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Accept %v\n", conn.RemoteAddr())

		go func() {
			req, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))
		}()
	}
}
