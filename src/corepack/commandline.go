package corepack

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func doCommandLine() {
	doMutex()
}

func doMutex() {
	m := new(sync.Mutex)
	
	for i:= 0; i<10; i++{
		go func(i int){
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}
	
	var input string
	fmt.Scanln(&input)
}

func doComandMain() {
	//Define flags
	maxp := flag.Int("max", 6, "The max value")
	//Parse
	flag.Parse()
	//Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
}
