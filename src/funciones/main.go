package funciones

import (
	"fmt"
)

func main() {
	e3()
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
