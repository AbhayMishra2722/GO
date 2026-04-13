package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Conversions")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your Rating", numRating+1)
	}

}
