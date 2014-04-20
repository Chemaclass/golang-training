package corepack

import (
	"fmt"
	"strings"
)

func main() {
	e1()
}

func e1() {
	arr := []byte("test")
	//	str := string([]byte{'t', 'e', 's', 't'})
	str := string(arr)
	fmt.Printf("arr: %s, str: %s", arr, str)
}

func doStrings() {
	fmt.Println(
		//true
		strings.Contains("test", "es"),
		//2
		strings.Count("test", "t"),
		// true
		strings.HasPrefix("test", "te"),
		// true
		strings.HasSuffix("test", "st"),
		// 1
		strings.Index("test", "e"),
		//"a-b"
		strings.Join([]string{"a", "b"}, "-"),
		// "aaaaa"
		strings.Repeat("a", 5),
		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),
		// []string{"a","b","c","d","e","f"}
		strings.Split("a-b-c-d-e-f", "-"),
		// "test"
		strings.ToLower("TEST"),
		// "TEST"
		strings.ToUpper("test"),
	)
}

func Main() {
	main()
}
