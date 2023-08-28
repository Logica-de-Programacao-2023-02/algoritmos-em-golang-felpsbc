package main

import "fmt"

func main() {

	var num, dobro, triplo, quadruplo int
	fmt.Println("Digite um número: ")
	fmt.Scan(&num)
	dobro = num + num
	fmt.Println("O dobro desse número é: ", dobro)
	triplo = num + num + num
	fmt.Println("O triplo desse número é: ", triplo)
	quadruplo = num + num + num + num
	fmt.Println("O quadruplo desse número é: ", quadruplo)

}
