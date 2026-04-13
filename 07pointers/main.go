package main

import "fmt"

func main() {

	fmt.Println("Pointers")

	var ptr *int
	fmt.Println("The Value of pointer is", ptr)

	var myNumber = 20
	var ptr1 = &myNumber                                 // & sign isused for reference
	fmt.Println("THe Value of second pointer is", ptr1)  //prints the address of myNumber
	fmt.Println("THe Value of second pointer is", *ptr1) //*ptr1 prints the actual value

	*ptr1 = *ptr1 + 5
	fmt.Println("New value is", myNumber)
}
