package listas

import (
	"fmt"
)

func main() {
	e5()
}

// Maps3
func e5() {
	e := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
	}
	if name, ok := e["N"]; ok {
		fmt.Println(name, ok)
	}
}

// Maps2
func e4() {
	e := make(map[string]string)
	e["H"] = "Hydrogen"
	e["He"] = "Helium"
	e["Li"] = "Lithium"
	e["Be"] = "Beryllium"
	e["B"] = "Boron"
	e["C"] = "Carbon"
	e["N"] = "Nitrogen"

	if name, ok := e["Un"]; ok {
		fmt.Println(name, ok)
	}
}

// Maps
func e3() {
	//var x map[string]int
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	delete(x, "key")
	fmt.Println(x)
}

// Copy slice
func e2() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

//append slice
func e1() {
	fmt.Println("listas.main")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func Main() {
	fmt.Println("listas.Main")
	main()
}
