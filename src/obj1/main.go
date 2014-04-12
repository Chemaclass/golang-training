package obj1

import (
	"fmt"
)

type Humano struct {
	nombre string
	edad   int
}

func (h Humano) String() string {
	return fmt.Sprintf("Humano{%s, %d}", h.nombre, h.edad)
}

func main() {
	fmt.Println("obj1.main BEGIN")
	fmt.Printf("%v\n", Humano{"Chema", 21})
	fmt.Println("obj1.main END")
}

func Main() {
	main()
}
