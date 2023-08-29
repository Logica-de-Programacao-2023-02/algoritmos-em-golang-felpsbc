package main

import "fmt"

func main() {

	var soma, div, i float64

	for {
		fmt.Println("Escreva um número: ")
		fmt.Scan(&i)
	}
	if i == 0 {
		break
	}
	soma += i;
	div++
}
