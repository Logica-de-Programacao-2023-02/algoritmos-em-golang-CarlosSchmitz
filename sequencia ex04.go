package main

import "fmt"

func main() {

	var num1, num2, num3 float32

	fmt.Print("Primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Terceiro número: ")
	fmt.Scan(&num3)

	var peso1, peso2, peso3 float32

	fmt.Print("Peso do primeiro: ")
	fmt.Scan(&peso1)

	fmt.Print("Peso do segundo: ")
	fmt.Scan(&peso2)

	fmt.Print("peso do terceiro: ")
	fmt.Scan(&peso3)

	res := (num1*peso1 + num2*peso2 + num3*peso3) / (peso1 + peso2 + peso3)

	fmt.Print("A média ponderada é: ", res)

}
