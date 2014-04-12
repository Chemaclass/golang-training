package params

import (
	"fmt"
	"strconv"
)

type Seq struct {
	nombre string
	act    int
	salto  int
}

func intSeq(p ...string) func() (string, int) {

	var act, salto int
	var nombre string

	switch len(p) {
	case 0:
		nombre, act, salto = "Seq", 0, 1
	case 1:
		nombre, act, salto = p[0], 0, 1
	case 2:
		t, _ := strconv.Atoi(p[1])
		nombre, act, salto = p[0], t, 1
	default:
		t, _ := strconv.Atoi(p[1])
		t2, _ := strconv.Atoi(p[2])
		nombre, act, salto = p[0], t, t2
	}

	seq := Seq{nombre, act, salto}

	return func() (string, int) {
		seq.act += seq.salto
		return seq.nombre, seq.act
	}
}

func main(params...interface{}) {
	nextInt := intSeq("Mi Seq1", "1", "2")

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq("Sec-2", "1")
	fmt.Println(newInts())

	newInts2 := intSeq()
	n, i := newInts2()
	fmt.Printf("-- %s %d", n, i)

}

func Main(params...interface{}){
	main(params)
}
