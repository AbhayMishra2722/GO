package main

import "fmt"

func main() {
	fmt.Println("Structs")
	abhay := User{"Abhay", "abhaymishra2722@gmail.com", true, 21}
	fmt.Println(abhay)
	fmt.Printf("Details are : %+v\n", abhay)

	fmt.Printf("Name is %v and Email is %v.", abhay.Name, abhay.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
