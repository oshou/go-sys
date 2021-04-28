package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	messages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil

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
		dump, err := httputil.DumpResponse(res, true)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dump))
		current++
		if current == len(sendMessages) {
			break
		}
	}
	conn.Close()
}
