package main

import "fmt"

func main() {
	fmt.Println("Maps")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("The List is :", languages)
	fmt.Println("JS stands for :", languages["JS"])

	delete(languages, "RB")
	fmt.Println("The List is :", languages)

}
