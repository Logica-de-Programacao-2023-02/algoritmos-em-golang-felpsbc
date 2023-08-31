package main

import "fmt"

func main() {
	var soma, div, i float64

	for {
		fmt.Println("Escreva um número: ")
		fmt.Scan(&i)

		if i == 0 {
			break
		}

		soma += i
		div++
	}

	if div > 0 {
		media := soma / div
		fmt.Printf("A média aritmética é: %.2f\n", media)
	} else {
		fmt.Println("não houve nenhum número inserido.")
	}
}
