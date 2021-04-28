package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// dial connection
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// create request
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		panic(err)
	}
	req.Write(conn)

	// read response
	res, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))
}
