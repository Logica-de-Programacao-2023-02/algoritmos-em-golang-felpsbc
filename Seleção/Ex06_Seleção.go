package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Println("Escreva dois números")
	fmt.Scan(&num1, &num2)
	resultado := num1 * num2
	soma := num1 + num2

	if num1 >= 0 && num2 >= 0 {
		fmt.Println("O resultado da multiplicação entre os números é: ", resultado)
	} else {
		fmt.Println("O resultado da soma dos números é: ", soma)
	}
}
