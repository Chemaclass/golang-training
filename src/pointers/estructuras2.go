package pointers

import "fmt"

type Persona struct {
	Nombre string
}

func (p *Persona) Saludar() {
	fmt.Println("Hola, me llamo", p.Nombre)
}

type Android struct {
	Persona
	Model string
}
