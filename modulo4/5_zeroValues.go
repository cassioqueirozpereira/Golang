package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// Se não é atribuído valor as variáveis na criação das mesmas
	// Elas já vem com valores iniciais
	fmt.Printf("Inteiro: %v\n", i)
	fmt.Printf("Float: %v\n", f)
	fmt.Printf("Bool: %v\n", b)
	fmt.Printf("String: %v\n", s)
}