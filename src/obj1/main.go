package obj1

import (
	"fmt"
)

func main() {
	fmt.Println("obj1.main BEGIN")
	h := Humano{"Chema", 21}
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", h.Trabajar())
	fmt.Println("obj1.main END")
}

func Main() {
	main()
}
