package main

import "fmt"

func main() {

	var peso, altura, imc, altura2 float64
	fmt.Print("Digite seu peso ?")
	fmt.Scanln(&peso)
	fmt.Print("Digite sua altura ?")
	fmt.Scan(&altura)
	altura2 = altura * altura
	imc = peso / altura2
	fmt.Println("O seu imc é: ", imc)

}
