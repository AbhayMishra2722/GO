package main

import "fmt"

const LoginToken string = "HBVYYVYVVYVVUY" //PUBLIC

func main() {
	var username string = "Abhay"
	fmt.Println(username)
	fmt.Printf("the type is %T \n", username)

	var isLoggedIN bool = true
	fmt.Println("isLoggedIN")
	fmt.Printf("the type is %T \n", isLoggedIN)

	var smallValue uint8 = 10
	fmt.Println(smallValue)
	fmt.Printf("the type is %T \n", smallValue)

	var smallFloat float32 = 21.2525
	fmt.Println(smallFloat)
	fmt.Printf("the type is %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("the type is %T \n", anotherVariable)

	//implicit type

	var name = "Abhay"
	fmt.Println(name)

	//no var style

	number := 20000
	fmt.Println(number)

	fmt.Println(LoginToken)
	fmt.Printf("the type is %T \n", LoginToken)

}
