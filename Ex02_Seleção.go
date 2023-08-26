package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Println("Escreva três números")
	fmt.Scan(&num1, &num2, &num3)

	if num1 < num2 && num1 < num3 {
		fmt.Println("O primeiro número é o menor = ", num1)
	} else if num2 < num1 && num2 < num3 {
		fmt.Println("O segundo número é o menor = ", num2)
	} else if num3 < num1 && num3 < num2 {
		fmt.Println("O terceiro número é o menor = ", num3)
	} else {
		fmt.Println("Os números são iguais", num1, num2, num3)
	}

}
