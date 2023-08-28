package main

import "fmt"

func main() {

	var num int
	fmt.Println("Escreva um número: ")
	fmt.Scan(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("O número é divisível por 3 e 5 ao mesmo tempo")
	} else {
		fmt.Println("O número não é divisível por 3 e 5 ao mesmo tempo")
	}

}
