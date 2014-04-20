package mtest

import "testing"
import "fmt"

type testPair struct {
	values  []float64
	average float64
}

var tests = []testPair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{1, 1}, 0}, //FAIL
	{[]float64{1, -1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			msg := fmt.Sprintln("For", pair.values,
				"expected", pair.average,
				"got", v)
			fmt.Print(msg)
			t.Error(msg)
		} else {
			fmt.Printf("ok\n")
		}
	}
}

func Average(v []float64) float64 {
	var t float64 = 0
	for _, v := range v {
		t += v
	}
	return t / float64(len(v))
}

func main() {
	TestAverage(new(testing.T))
}

func Main() {
	fmt.Println("test.Main")
	main()
 }
