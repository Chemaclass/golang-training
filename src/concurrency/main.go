package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	e3()
}

//Select
func e3() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("<-c1|", msg1)
			case msg2 := <-c2:
				fmt.Println("<-c2|", msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

//Pin-pon
func e2() {
	pinger := func(c chan<- string) {
		for i := 0; ; i++ {
			c <- "ping"
		}
	}
	ponger := func(c chan<- string) {
		for i := 0; ; i++ {
			c <- "pong"
		}
	}
	printer := func(c <-chan string) {
		for {
			msg := <-c
			fmt.Println(msg)
			time.Sleep(time.Second * 1)
		}
	}

	c := make(chan string)
	go pinger(c)
	go ponger(c)
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
