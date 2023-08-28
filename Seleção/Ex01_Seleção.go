package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Println("Escreva dois números:")
	fmt.Scan(&num1, &num2)

	if num1 > num2 {
		fmt.Println("Número 1 é maior que o número 2")
	} else if num1 == num2 {
		fmt.Println("Número 1 é igual ao número 2")
	} else {
		fmt.Println("Número 1 é menor que número 2")
	}

}
