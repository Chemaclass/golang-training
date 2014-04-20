package corepack

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	doIOCorto()
}


func doIOCorto() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string (bs)
	fmt.Println(str)
}

func doIO() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error a> ", err)
		//handle the error here
		return
	}
	defer file.Close()

	//get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("error b> ", err)
		return
	}

	//read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("error c> ", err)
		return
	}

	str := string(bs)
	fmt.Println(str)
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
	fmt.Println("corepack.Main()")
	main()
}
