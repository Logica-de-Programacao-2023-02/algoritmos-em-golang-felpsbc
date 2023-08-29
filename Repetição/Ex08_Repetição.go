package main

import "fmt"

func main() {

	var num int
	fmt.Println("Escreva um número: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println("Os divisores de", num, "são", i)
		}
	}

}
