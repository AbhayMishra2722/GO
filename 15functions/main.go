package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greeter()

	result := adder(3, 4)
	fmt.Println("Sum is :", result)

	proResult := proAdder(2, 4, 6, 8, 9, 11)
	fmt.Println("Pro Result is :", proResult)
}

func greeter() {
	fmt.Println("Hello")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(value ...int) int {
	sum := 0
	for _, val := range value {
		sum += val
	}
	return sum
}
