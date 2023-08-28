package main

import "fmt"

func main() {

	var peso float64
	fmt.Println("Escreva seu peso(kg)")
	fmt.Scan(&peso)

	peso = peso * 2.2
	fmt.Println("O peso em librar e: ", peso)

}
