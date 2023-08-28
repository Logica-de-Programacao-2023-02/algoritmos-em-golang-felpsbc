package main

import "fmt"

func main() {

	var trabalho, diaria int

	fmt.Println("Quantos dia você tabalhou ? ")
	fmt.Scan(&trabalho)
	fmt.Println("Qual valor da sua diária ? ")
	fmt.Scan(&diaria)

	salario := diaria * trabalho
	fmt.Println("O seu salário deve ser: ", salario)

}
