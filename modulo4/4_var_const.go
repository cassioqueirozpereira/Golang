package main

import "fmt"

func main() {
	// Váriáveis
	// var + nome da variável + tipo
	var nome string
	nome = "Victor"
	fmt.Println(nome)

	nome = "Cássio"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Println(idade)

	// var + nome da variável, o go já entende
	var name = "Cássio"
	fmt.Println(name)
	
	var a = true
	fmt.Println(a)

	// declarando mais de uma variável do mesmo tipo
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Maneira mais simples e a mais usada em Golang, para declarar uma variável
	f := "Linux"
	fmt.Println(f)

	// Constantes
	const idadeVictor = 14
	fmt.Println(idadeVictor)

	
}