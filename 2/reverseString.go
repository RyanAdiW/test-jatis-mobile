package main

import "fmt"

func main() {
	input := "jatis"
	output := ""

	rnsInput := []rune(input)
	for i, j := 0, len(rnsInput)-1; i < j; i, j = i+1, j-1 {
		rnsInput[i], rnsInput[j] = rnsInput[j], rnsInput[i]
	}
	output = string(rnsInput)

	fmt.Println("Input = ", input)
	fmt.Println("Output = ", output)
}
