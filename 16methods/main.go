package main

import "fmt"

func main() {
	abhay := User{"Abhay", "abhaymishra2722@gmail.com", true, 21}
	//fmt.Println(abhay)
	//fmt.Printf("Details are : %+v\n", abhay)
	//fmt.Printf("Name is %v and Email is %v.", abhay.Name, abhay.Email)

	abhay.GetStatus()
	abhay.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // this is the method to make a method
	fmt.Println("Is User Active", u.Status)

}

func (u User) NewEmail() {
	u.Email = "asdf@gmail.con"
	fmt.Println("New Email :", u.Email)
}
