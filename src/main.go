package main

import (
	"commands"
	"even"
	"fibonacci"
	"flag"
	"fmt"
	"obj1"
	"params"
	"sum"
)

func main() {
	doScanExcample()
}

func doScanExcample() {
	var name string
	fmt.Print("What's your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s\n", name)
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
	fibonacci.Main()
}
