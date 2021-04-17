package main

import "fmt"

type Talker interface {
	Talk()
}

type Greeter struct {
	name string
}

var _ Talker = Greeter{}

func (g Greeter) Talk() {
	fmt.Printf("Hello, my name is %s\n", g.name)
}

func main() {
	var talker Talker
	talker = &Greeter{"nanana"}
	talker.Talk()
}
