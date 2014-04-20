package pointers

import (
	"fmt"
)

func main() {
	e2()
}

//Struct and Interfaces
func e2() {
	r := Rectangulo{1, 2, 3, 4}
	fmt.Printf("%s", r)
	c := Circulo{1.23, 4.56, 7.89}
	fmt.Printf("%s", c)
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
	pzero(&x) // sí modifica x
	fmt.Printf("x es %d\n", x)
}

func Main() {
	fmt.Println("pointers.Main")
	main()
}
