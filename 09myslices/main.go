package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Mango", "Banana"}
	fmt.Println("The Fruit List is :", fruitList)

	fruitList = append(fruitList, "Kiwi", "Grapes")
	fmt.Println("The Updated Fruit List is :", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Sliced Fruit List is :", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 250
	highScores[1] = 563
	highScores[2] = 365
	highScores[3] = 486
	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
