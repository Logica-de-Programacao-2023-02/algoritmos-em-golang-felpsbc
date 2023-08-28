package main

import "fmt"

func main() {

	var produto float64
	fmt.Println("Escreva o valor do produto: ")
	fmt.Scan(&produto)

	produto = produto - produto*0.10
	fmt.Println("O valor do produto com desconto e :", produto)

}
