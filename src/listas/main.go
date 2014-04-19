package listas

import (
	"fmt"
)

func main() {
	e2()
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
