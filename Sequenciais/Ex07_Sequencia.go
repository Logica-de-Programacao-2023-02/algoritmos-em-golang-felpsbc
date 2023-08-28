package main

import "fmt"

func main() {

	var num int
	fmt.Println("escreva um numero: ")
	fmt.Scan(&num)

	antecessor := num - 1
	sucessor := num + 1

	fmt.Println("O sucessor e: ", sucessor)
	fmt.Println("O antecessor e: ", antecessor)

}
