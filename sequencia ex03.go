package main

import "fmt"

func main() {

	var peso, altura int

	fmt.Print("Qual é seu peso? (em kg) ")
	fmt.Scan(&peso)

	fmt.Print("Qual é sua alura (em m) ")
	fmt.Scan(&altura)

	IMC := peso / (altura ^ 2)

	fmt.Print("seu IMC é ", IMC)

}
