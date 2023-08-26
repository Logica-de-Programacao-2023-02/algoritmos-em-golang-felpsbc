package main

import "fmt"

func main() {

	var idade, DiasVida int

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)
	DiasVida = idade * 365
	fmt.Print("Você já viveu: ", DiasVida, " dias.")

}
