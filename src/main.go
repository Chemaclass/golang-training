package main

import (
	"even"
	"flag"
	"fmt"
	"params"
	"sum"
)

func main() {

	params.Main()
	even.Main(1, 2, 72)
	sum.Main(1, 2, 3, 4, 5)
	//Args
	flag.Parse()
	a := flag.Args()
	for _, i := range a {
		fmt.Printf("> %s\n", i)
	}
	fmt.Printf(">> %s", flag.Arg(0))
	//Args end

}
