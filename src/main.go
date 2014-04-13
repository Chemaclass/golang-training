package main

import (
	"commands"
	"even"
	"flag"
	"fmt"
	"obj1"
	"params"
	"sum"
	"fibonacci"
)

func main() {
	fibonacci.Main()
}

func act1() {
	params.Main()
	even.Main(1, 2, 72)
	sum.Main(1, 2, 3, 4, 5)
	//Args
	flag.Parse()
	a := flag.Args()
	for _, i := range a {
		fmt.Printf("> %s\n", i)
	}
	//Args end
	obj1.Main()
	commands.Main()
}
