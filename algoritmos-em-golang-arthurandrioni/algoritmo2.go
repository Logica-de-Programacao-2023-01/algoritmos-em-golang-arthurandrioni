package main

import "fmt"

func main() {
	num1 := 0
	num2 := 0
	num3 := 0

	fmt.Print("Qual o primeiro numero? ")
	fmt.Scan(&num1)
	fmt.Print("Qual o segundo numero? ")
	fmt.Scan(&num2)
	fmt.Print("Qual o terceiro numero? ")
	fmt.Scan(&num3)

	if num1 > num2 {
		fmt.Print("O menor numero é: ", num2)
	} else if num1 > num3 {
		fmt.Print("O menor numero é: ", num3)
	} else {
		fmt.Print("O menor numero é: ", num1)

	}
}
