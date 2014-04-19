package funciones

import (
	"fmt"
)

func main() {
	e1()

}

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
