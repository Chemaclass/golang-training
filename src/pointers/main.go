package pointers

import (
	"fmt"
)

func main() {
	e4()
}

func e4() {
	c := &Circulo{1, 2, 3}
	r := &Rectangulo{1, 2}
	l := new(MultiFormas)
	// Refactorizo y creo el método Add()
	// l.Formas = append(l.Formas, c, r)
	l.Add(c, r)
	a := l.Area()
	fmt.Println("area total: ", a)
}

func e3() {
	a := new(Android)
	a.Nombre = "Chema"
	a.Saludar()
}

//Struct and Interfaces
func e2() {
	r := Rectangulo{alto: 10, ancho: 4}
	fmt.Printf("%s", r)
	fmt.Printf("Area del rect %.2f\n", r.Area())
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
