package main

import "fmt"

func main() {

	var numero, anterior, sucessor int

	fmt.Println("Escreva um número: ")
	fmt.Scan(&numero)

	anterior = numero - 1
	sucessor = numero + 1

	fmt.Println("O número anterior é :", anterior)
	fmt.Println("O número sucessor é :", sucessor)

}
