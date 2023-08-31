package main

import (
	"fmt"
)

func main() {
	var num, maior int

	maior = 0

	fmt.Println("Digite vários números inteiros para que mostre o maior e Digite 0 para parar.")

	for {
		fmt.Print("Digite um número: ")
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		if num > maior {
			maior = num
		}
	}

	if maior == 0 {
		fmt.Println("Nenhum número foi digitado.")
	} else {
		fmt.Printf("O maior número digitado foi: %d\n", maior)
	}
}
