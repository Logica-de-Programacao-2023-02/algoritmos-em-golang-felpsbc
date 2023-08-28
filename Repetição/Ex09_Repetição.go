package main

import "fmt"

func main() {

	var preco float64
	fmt.Println("Qual o preço do produto ?")
	fmt.Scan(&preco)

	preco = preco - preco*0.10
	fmt.Println("O preço da roupa com descontro é: ", preco, "reais ")

}
