package main

import "fmt"

func main() {

	var peso, altura, imc, altura2 float64
	fmt.Println("Digite seu peso ?")
	fmt.Scanln(&peso)
	fmt.Println("Digite sua altura ?")
	fmt.Scan(&altura)
	altura2 = altura * altura
	imc = peso / altura2
	fmt.Println("Seu imc é:", imc)

	if imc < 18.5 {
		fmt.Println("Você está abaixo do peso")
	} else if imc < 24.9 && imc > 18.5 {
		fmt.Println("Você está no peso ideal")
	} else {
		fmt.Println("Você está acima do peso")
	}
}
