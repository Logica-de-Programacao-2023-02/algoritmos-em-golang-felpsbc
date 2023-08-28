package main

import "fmt"

func main() {

	var num1, num2, num3 int
	fmt.Println("Digite Três números:")
	fmt.Scan(&num1, &num2, &num3)

	if num1 <= num2 && num1 <= num3 && num2 <= num3 {
		fmt.Println("Os números na ordem crescente são: ", num1, num2, num3)
	} else if num1 <= num2 && num1 <= num3 && num2 >= num3 {
		fmt.Println("Os números na ordem crescente são: ", num1, num3, num2)
	} else if num1 >= num2 && num1 <= num3 && num2 <= num3 {
		fmt.Println("Os números na ordem crescente são: ", num2, num1, num3)
	} else if num1 >= num2 && num1 >= num3 && num2 <= num3 {
		fmt.Println("Os números na ordem crescente são: ", num2, num3, num1)
	} else if num1 >= num2 && num1 >= num3 && num2 <= num3 {
		fmt.Println("Os números na ordem crescente são: ", num3, num1, num2)
	} else if num1 >= num2 && num1 >= num3 && num2 >= num3 {
		fmt.Println("Os números na ordem crescente são: ", num3, num2, num1)
	}
}
