package main

import "fmt"

func main() {

	var salario float64

	fmt.Print("Qual seu salário ? ")
	fmt.Scan(&salario)
	salario = salario + 0.15*salario
	fmt.Println("O seu salario atual é de ", salario)

}
