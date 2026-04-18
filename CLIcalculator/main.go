package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var num1, num2 float64
		var operator string

		fmt.Println("\n--- CLI Calculator ---")
		fmt.Println("Enter expression (e.g. 5 + 3) or type Bye to quit:")

		input, _ := reader.ReadString('\n')

		if input == "Bye\n" {
			fmt.Println("Goodbye!")
			break
		}

		_, err := fmt.Sscanf(input, "%f %s %f", &num1, &operator, &num2)
		if err != nil {
			fmt.Println("Invalid input format. Use: number operator number")
			continue
		}

		switch operator {
		case "+":
			fmt.Println("Result:", num1+num2)
		case "-":
			fmt.Println("Result:", num1-num2)
		case "*":
			fmt.Println("Result:", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero")
			} else {
				fmt.Println("Result:", num1/num2)
			}
		default:
			fmt.Println("Invalid operator. Use +, -, *, /")
		}
	}
}
