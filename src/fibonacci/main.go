package fibonacci

import (
	"fmt"
)

func fibonacci(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

func printFibonacci(n int) {
	for _, term := range fibonacci(10) {
		fmt.Printf("%v - ", term)
	}
}

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []int{1,2,3}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", Map(f,m))
}

func Main() {
	main()
}
