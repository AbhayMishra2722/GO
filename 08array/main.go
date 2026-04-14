package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Kiwi"
	fmt.Println("FruitList is :", fruitList)
	fmt.Println("Length of FruitList is :", len(fruitList))

	var vegList = [3]string{"Beans", "Potato", "Onion"}
	fmt.Println("Vegetable List is :", vegList)
	fmt.Println("Length of Vegetable List is :", len(vegList))
}
