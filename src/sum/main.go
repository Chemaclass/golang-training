package sum

import (
	"fmt"
)

type Vec struct {
	list []int
}

func (v Vec) Sum() (s int) {
	for _, value := range v.list {
		s += value
	}
	return
}

func (v Vec) String() (r string) {
	r = "Vec {"
	for _, v := range v.list {
		r += fmt.Sprintf("%d,", v)
	}
	r += fmt.Sprintf("}")
	return
}

func Main(p ...int) {
	fmt.Println("sum BEGIN")
	v := Vec{p}
	fmt.Printf("%s\n", v)
	fmt.Printf("La suma es: %d\n", v.Sum())
	fmt.Println("sum END")
}
