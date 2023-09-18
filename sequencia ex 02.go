package main

import "fmt"

func main() {

	var num int

	fmt.Print("Número: ")
	fmt.Scan(&num)

	fmt.Print("O número escolhido foi ", num, " seu dobro é ", num*2, " seu triplo é ", num*3, " e seu quádruplo é ", num*4)
}
