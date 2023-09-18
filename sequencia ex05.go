package main

import "fmt"

func main() {

	var idade float64

	fmt.Print("Sua idade em anos: ")
	fmt.Scan(&idade)

	fmt.Print("Sua idade em anos é: ", idade, ", em meses é: ", idade*12, ", e em dias é: ", idade*365)

}
