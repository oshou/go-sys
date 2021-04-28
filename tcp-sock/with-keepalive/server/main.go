package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
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
			defer conn.Close()

			for {
				conn.SetReadDeadline(time.Now().Add(5 * time.Second))

				// read request
				req, err := http.ReadRequest(bufio.NewReader(conn))
				if err != nil {
					netErr, ok := err.(net.Error)
					if ok && netErr.Timeout() {
						fmt.Println("Timeout")
						break
					} else if err == io.EOF {
						break
					}
					panic(err)
				}
				dump, err := httputil.DumpRequest(req, true)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(dump))
				content := "Hello World\n"

				// write response
				res := http.Response{
					StatusCode:    200,
					ProtoMajor:    1,
					ProtoMinor:    1,
					ContentLength: int64(len(content)),
					Body:          ioutil.NopCloser(strings.NewReader(content)),
				}
				res.Write(conn)
			}
		}()
	}
}
