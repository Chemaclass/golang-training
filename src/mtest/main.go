package mtest

import "testing"
import "fmt"

type testPair struct {
	values    []float64
	resultado float64
}

type testTipo struct {
	tipo     string
	testPair []testPair
}

var tests = []testTipo{
	testTipo{"media", testsMedia},
	testTipo{"max", testsMax},
	testTipo{"min", testsMin},
}

var testsMedia = []testPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{1, 1}, 0}, //FAIL
	{[]float64{1, -1}, 0},
	{[]float64{}, 0},
}
var testsMax = []testPair{
	{[]float64{1, 2}, 2},
	{[]float64{5, 4}, 5},
	{[]float64{3, 7}, 3}, // FAIL
}
var testsMin = []testPair{
	{[]float64{1, 2}, 1},
	{[]float64{5, 4}, 4},
	{[]float64{3, 7}, 7}, // FAIL
}

//Test Media
func Tests(t *testing.T) {
	for _, testsList := range tests {
		for _, pair := range testsList.testPair {

			var v float64
			switch testsList.tipo {
			case "media":
				v = Media(pair.values)
			case "max":
				v = Max(pair.values)
			case "min":
				v = Min(pair.values)
			}
			if v != pair.resultado {
				msg := fmt.Sprintln(testsList.tipo, "For", pair.values,
					"expected", pair.resultado,
					"got", v)
				fmt.Print(msg)
				t.Error(msg)
			} else {
				fmt.Printf("ok\n")
			}
		}
	}
}

//Devuelve la media de una lista
func Media(v []float64) float64 {
	if len(v) <= 0 {
		return 0
	}
	var t float64 = 0
	for _, v := range v {
		t += v
	}
	return t / float64(len(v))
}

//Devuelve el máximo de una lista
func Max(lista []float64) (m float64) {
	if len(lista) <= 0 {
		return 0
	}
	m = lista[0]
	for _, v := range lista {
		if v > m {
			m = v
		}
	}
	return
}

//Devuelve el menor de una lista
func Min(lista []float64) (m float64) {
	if len(lista) <= 0 {
		return 0
	}
	m = lista[0]
	for _, v := range lista {
		if v < m {
			m = v
		}
	}
	return
}

func main() {
	t := new(testing.T)
	Tests(t)
}

func Main() {
	fmt.Println("test.Main")
	main()
}
