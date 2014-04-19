package funciones

import (
	"fmt"
)

func main() {
	e6()
}

// Creamos un generador de números impares
func e6() {
	generadorImpares := func() func() int {
		i := 1
		return func() (r int) {
			r = i
			i += 2
			return
		}
	}
	siguienteImpar := generadorImpares()
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
	fmt.Println(siguienteImpar())
}

//Defer, Panic y Recover
func e5() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

//Defer
func e4() {
	f := func() {
		fmt.Println("1st")
	}
	s := func() {
		fmt.Println("2nd")
	}
	defer s()
	f()
}

//Recursion
func e3() {
	fmt.Printf("factorial:  %d", factorial(8))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

//Closure
func e2() {
	siguientePar := generadorPares()
	fmt.Println(siguientePar()) //0
	fmt.Println(siguientePar()) //2
	fmt.Println(siguientePar()) //4
}

func generadorPares() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//Calcular la media
func e1() {
	media := func(xs []float64) float64 {
		total := 0.0
		for _, v := range xs {
			total += v
		}
		return total / float64(len(xs))
	}

	xs := []float64{98, 93, 77, 82, 83}
	fmt.Printf("Media %.3f", media(xs))
}

func Main() {
	fmt.Println("funciones.Main")
	main()
}
