package main

import "fmt"

func main() {

	var diaria, dias int
	fmt.Println("digite sua diaria e seus dias trabalhados: ")
	fmt.Scan(&diaria, &dias)

	salario := diaria * dias
	fmt.Println("o valor do seu salario e :", salario)

}
