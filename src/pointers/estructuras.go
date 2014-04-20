package pointers

import (
	"fmt"
	"math"
)

type MultiFormas struct {
	Formas []Forma
}

func (m *MultiFormas) Add(a ...Forma) {
	for _, v := range a {
		m.Formas = append(m.Formas, v)
	}
}

func (m *MultiFormas) Area() float64 {
	var area float64
	for _, s := range m.Formas {
		area += s.Area()
	}
	return area
}

type Forma interface {
	Area() float64
}

type Rectangulo struct {
	alto, ancho float64
}

func (self Rectangulo) String() string {
	return fmt.Sprintf("Rect{ alto:%.2f, ancho:%.2f}\n", self.alto, self.ancho)
}

func (r *Rectangulo) Area() float64 {
	return r.alto * r.ancho
}

type Circulo struct {
	p1, p2, r float64
}

func (c *Circulo) Area() float64 {
	return math.Pi * (c.r * c.r)
}
