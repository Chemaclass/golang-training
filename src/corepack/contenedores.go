package corepack

import (
	"container/list"
	"fmt"
	"sort"
)

func doContenedores() {
	doOrdenar()
}

func doOrdenar() {
	kids := []Person{
		{"Jill", 9},
		{"Chema", 20},
		{"Jack", 11},
	}
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

func doList() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

}
