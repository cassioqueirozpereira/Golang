package main

import "fmt"

func main() {
	// %T é usado para representar o tipo da variável
	fmt.Printf("Type: %T - Value: %v \n", true, true)
	// string - sequência de bytes
	fmt.Printf("Type: %T - Value: %v \n", "Cássio", "Cássio")
	// Número entre aspas é convertido para string
	fmt.Printf("Type: %T - Value: %v \n", "1", "1")
	// %v é usado para representar o valor de todos os tipos de variáveis
	fmt.Printf("Type: %T - Value: %v \n", 1, 1)
	// 
	fmt.Printf("Type: %T - Value: %v \n", 1.234, 1.234)
}