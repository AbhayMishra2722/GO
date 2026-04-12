package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var name string = "Welcome"
	fmt.Println(name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Anything")

	// ,ok or err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Everything", input)
	fmt.Printf("the type of output is %T", input)

}
