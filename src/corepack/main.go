package corepack

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"errors"
)

func main() {
	doContenedores()
}

func doErrors(){
	err := errors.New("error message")
	err.Error()
}

//Recorre toda la estructura de forma recursiva
func doLeerConWalk() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}

//Leer un directorio
func doLeerDir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func doCreateFile() {
	file, err := os.Create("../test.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("test\n")
}

func doIOCorto() {
	bs, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		return
	}
	str := string(bs)
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
