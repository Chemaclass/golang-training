package even

import (
	"fmt"
)

func Even(i int) bool {
	return i%2 == 0
}

func odd(i int) bool {
	return i%2 != 0
}

func WhatIs(i int) string {
	if odd(i) {
		return "Odd"
	}
	return "Even"
}

func main(params []int) {
	for _, n := range params {
		fmt.Printf("%.2d is %s\n", n, WhatIs(n))
	}
}

func Main(params ...int) {
	fmt.Println("even.Main BEGIN")
	main(params)
	fmt.Println("even.Main END")
}
