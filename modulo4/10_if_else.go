package main

import "fmt"

// IF / ELSE

func main() {

	valor := 7
	// if (condição) { <ação> }
	if valor == 1 {
		fmt.Println("\nValor é igual a 1\n")
	} else if valor == 0 {
		fmt.Println("\nValor é igual a 0\n")
	} else {
		fmt.Println("\nValor é diferente de 0 e 1\n")
	}

	if valor % 2 == 0 {
		fmt.Printf("%d é par\n", valor)
	} else {
		fmt.Printf("%d é ímpar\n", valor)
	}
}