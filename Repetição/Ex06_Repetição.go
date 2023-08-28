package main

import "fmt"

func main() {

	var num int
	fmt.Println("Escreva o numero que deseja mulitplicar :")
	fmt.Scan(&num)

	for mult := 1; mult <= 10; mult++ {
		fmt.Println("O resultado da sua tabuada e", num*mult)
	}
}
