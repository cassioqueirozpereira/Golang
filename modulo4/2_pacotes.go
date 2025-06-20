package main

import (
	"fmt"
	// Renomeando o nome do pacote strings para str
	str "strings"
)

func main() {
	fmt.Println("Hello, world!")

	// Utilizando o pacote renomeado
	fmt.Println(str.Split("Cássio", ""))
}