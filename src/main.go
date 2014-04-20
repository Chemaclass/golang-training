package main

import (
	"concurrency"
	"fmt"
	"funciones"
	"listas"
	"mtest"
	"pointers"
)

func main() {
	doTest()
}

func doTest() {
	mtest.Main()
}

func doConcurrency() {
	concurrency.Main()
}
func doPointers() {
	pointers.Main()
}

func doFunciones() {
	funciones.Main()
}

func doListas() {
	listas.Main()
}

func doScanExcample() {
	var name string
	fmt.Print("What's your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s\n", name)
}
