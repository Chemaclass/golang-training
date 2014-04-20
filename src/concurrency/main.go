package concurrency

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	e1()
}

func e1() {
	for i:= 0; i< 10; i++{
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
