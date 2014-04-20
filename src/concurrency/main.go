package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	e2()
}

func e2() {
	pinger := func(c chan string) {
		for i := 0; ; i++ {
			c <- "ping"
		}
	}
	printer := func(c chan string) {
		for {
			msg := <-c
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}

	c := make(chan string)
	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}

func e1() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func Main() {
	fmt.Println("concurrency.Main")
	main()
}
