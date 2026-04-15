package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	//}

	//for i := range days {
	//fmt.Println(days[i])
	//}

	//for i, days := range days {
	//	fmt.Printf("Index is %v and the Day is %v \n", i, days)
	//}

	c := 1
	for c <= 10 {

		if c == 2 {
			goto d
		}

		if c == 5 {
			break
		}
		fmt.Println(c)
		c++
	}
d:
	fmt.Println("Abhay")
}
