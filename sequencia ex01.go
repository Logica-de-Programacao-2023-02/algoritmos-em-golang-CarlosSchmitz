package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Print("Primeiro número: ")
	fmt.Scan(&num1)
	fmt.Print("Seundo número: ")
	fmt.Scan(&num2)
	fmt.Print("Terceiro número ")
	fmt.Scan(&num3)

	res := num1 + num2 + num3

	fmt.Println("A soma dos números é: ", res)
}
