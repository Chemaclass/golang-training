package pointers

import (
	"fmt"
)

func main() {
	e1()
}

func e1() {
	zero := func(p int) {
		p = 0
	}
	pzero := func(p *int) {
		*p = 0
	}
	x := 5
	zero(x) // No modifica x
	fmt.Printf("x es %d|", x)
	pzero(&x)// sí modifica x
	fmt.Printf("x es %d\n", x)
}

func Main() {
	fmt.Println("pointers.Main")
	main()
}
