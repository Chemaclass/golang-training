package pointers

import "fmt"

type Rectangulo struct {
	  alto, ancho float64
}

func (self Rectangulo) String() string {
	return fmt.Sprintf("Rect{ alto:%.2f, ancho:%.2f}\n", self.alto,self.ancho)
}

func (r *Rectangulo) Area() float64 {	
	return r.alto * r.ancho
}

type Circulo struct {
	x, y, z float64
}

func (self Circulo) String() string {
	return fmt.Sprintf("Cir{ x:%.2f, y:%.2f, z:%.2f}\n", self.x, self.y, self.z)
}
