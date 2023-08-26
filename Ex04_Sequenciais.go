package main

import "fmt"

func main() {

	var num1, num2, num3, peso1, peso2, peso3, np1, np2, np3 int
	var media int

	fmt.Println("Digite o primeiro número: ")
	fmt.Scan(&num1)
	fmt.Println("Digite o segundo número: ")
	fmt.Scan(&num2)
	fmt.Println("Digite o terceiro número: ")
	fmt.Scan(&num3)
	fmt.Println("Digite o primeiro peso: ")
	fmt.Scan(&peso1)
	fmt.Println("Digite o segundo peso: ")
	fmt.Scan(&peso2)
	fmt.Println("Digite o terceiro peso: ")
	fmt.Scan(&peso3)
	np1 = num1 * peso1
	np2 = num2 * peso2
	np3 = num3 * peso3
	media = (np1 + np2 + np3) / (peso1 + peso2 + peso3)
	fmt.Println("A média ponderada é: ", media)

}
