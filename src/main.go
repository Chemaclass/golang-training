package main

import (
	"fmt"
	"funciones"
	"listas"
)

func main() {
	doFunciones()
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
