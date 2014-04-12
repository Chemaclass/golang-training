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

func (h Humano) Trabajar() string {
	return fmt.Sprintf("%s está trabajando", h.nombre)
}