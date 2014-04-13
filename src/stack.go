package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
	P(*s)
}

func (s *Stack) Pop(ret int) {
	s.i--
	ret = s.data[s.i]
	P(*s)
}

func (s *Stack) String() (str string) {
	for i := 0; i < s.i; i++ {
		str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func P(s Stack) {
	fmt.Printf("%v\n", s)
}

func main() {
	fmt.Println("stack begin")
	s := new(Stack)
	s.Push(1)
	s.Push(4)
	s.Push(3)
	fmt.Printf("%s\n", s)
	fmt.Println("stack end")
}
