package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "This needs to go in my file"
	file, err := os.Create("./mylcogofile.txt")
	if err != nil { // func checkNilError(err  error){ if error!= nil{panic(err)}} and then call this function where ever to check error
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is :", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}
func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data in the file is \n", data)
	fmt.Println("Text data in the file is \n", string(data))
}
