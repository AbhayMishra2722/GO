package main

import "fmt"

func main() {
	fmt.Println("IfElse")

	loginCount := 25
	var result string

	if loginCount > 10 {
		result = "Regular User"
	} else {
		result = "Not a Regular User"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("More")
	}
}
