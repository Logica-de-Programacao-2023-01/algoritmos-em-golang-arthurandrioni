package main

import (
	"fmt"
)

func main() {
	Num_1 := 0
	Num_2 := 0

	fmt.Print("Qual o primeiro numero? ")
	fmt.Scan(&Num_1)
	fmt.Print("Qual o segundo numero? ")
	fmt.Scan(&Num_2)

	if Num_1 > Num_2 {
		fmt.Print("O maior numero é: ", Num_1)

	} else {
		fmt.Print("O maior numero é: ", Num_2)
	}

}
