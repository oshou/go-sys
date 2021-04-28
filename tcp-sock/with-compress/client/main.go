package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	var (
		messages = []string{
			"ASCII",
			"PROGRAMMING",
			"PLUS",
		}
		current int      = 0
		conn    net.Conn = nil
	)

	for {
		// dial connection
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8080")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %d\n", current)
		}

		// create request
		req, err := http.NewRequest(
			"POST",
			"http://localhost:8080",
			strings.NewReader(messages[current]),
		)
		if err != nil {
			panic(err)
		}
		req.Header.Set("Accept-Encoding", "gzip")
		if err = req.Write(conn); err != nil {
			panic(err)
		}

		// read response
		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}
		defer res.Body.Close()

		if res.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(res.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, res.Body)
		}

		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		current++
		if current == len(messages) {
			break
		}
	}
}
