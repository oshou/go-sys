package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "ascii.jp:80")
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n"))

	res, err := http.ReadResponse(bufio.NewReader(conn), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Header)
	defer res.Body.Close()

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		panic(err)
	}
}
