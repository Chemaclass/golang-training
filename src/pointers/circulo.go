package pointers

import "fmt"

type Rectangulo struct {
	x1, x2, y1, y2 float64
}

func (self Rectangulo) String() string {
	return fmt.Sprintf("Rect{ x1:%.2f, x2:%.2f, y1:%.2f, y2:%.2f}\n", self.x1, self.x2, self.y1, self.y2)
}

type Circulo struct {
	x, y, z float64
}

func (self Circulo) String() string {
	return fmt.Sprintf("Cir{ x:%.2f, y:%.2f, z:%.2f}\n", self.x, self.y, self.z)
}
