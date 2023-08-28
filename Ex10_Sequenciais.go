package main

import "fmt"

func main() {

	var peso float64

	fmt.Println("Qual seu peso em kg? ")
	fmt.Scan(&peso)
	peso = peso * 2.20

	fmt.Println("Seu peso em Libras é : ", peso)

}
