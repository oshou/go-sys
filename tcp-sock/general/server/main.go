package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
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
			// read request
			req, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dump))

			// write response
			res := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello, World\n")),
			}
			res.Write(conn)
			defer conn.Close()
		}()
	}
}
