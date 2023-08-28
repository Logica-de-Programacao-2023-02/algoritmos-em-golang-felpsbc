package main

import "fmt"

func main() {

	var salario float64
	fmt.Println("Qual é o seu salário ? ")
	fmt.Scan(&salario)
	aumento1 := salario + salario*0.10
	aumento2 := salario + salario*0.5

	if salario >= 1000 {
		fmt.Println("O seu aumento irá ser de 5%. Esse vai ser seu novo salário: ", aumento2)
	} else {
		fmt.Println("O seu aumento irá ser de 10%. Esse vai ser seu novo salário: ", aumento1)
	}

}
