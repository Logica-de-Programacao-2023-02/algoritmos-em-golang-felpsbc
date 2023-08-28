package main

import "fmt"

func main() {

	var num int
	fmt.Println("Escreva um número entre 1 e 7:")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Println("Hoje é domingo")
	} else if num == 2 {
		fmt.Println("Hoje é segunda-feira")
	} else if num == 3 {
		fmt.Println("Hoje é terça-feira")
	} else if num == 4 {
		fmt.Println("Hoje é quarta-feira")
	} else if num == 5 {
		fmt.Println("Hoje é quinta-feira")
	} else if num == 6 {
		fmt.Println("Hoje é sexta-feira")
	} else if num == 7 {
		fmt.Println("Hoje é sabado")
	} else {
		fmt.Println("Esse número não corresponde a um dia da semana")
	}

}
