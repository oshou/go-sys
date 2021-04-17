package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `aaa
bbbb
ccccc`

func main() {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		fmt.Printf("%#v\n", scanner.Text())
	}
}
