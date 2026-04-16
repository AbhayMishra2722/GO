package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two") // last in first out
	fmt.Println("Hello")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // 0 1 2 3 4 so printing will happen reverse due to defer 4 3 2 1 0
	}
}
