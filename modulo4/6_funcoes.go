package main

import "fmt"

func main() {
	soma := soma(42, 13)
	fmt.Println(soma)

	sub := subtracao(10, 5)
	fmt.Println(sub)

	// O underline procegue mesmo sem a variável.
	nome, nome2, _ := printName("Cássio")
	fmt.Println(nome, nome2)

	nome, sobrenome := printFullName("Cássio", "Pereira")
	fmt.Println(nome, sobrenome)
}

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func printName(nome string) (string, string, string) {
	return nome, nome, nome
}

// Função começando com letra minúscula:
// Função PRIVADA
// Só pode ser usada no próprio pacote
func printFullName(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

// Função começando com letra maiúscula:
// Função PÚBLICA
// Função pode ser usada em outros pacotes
// Como utilizar:
// main.PrintSobrenome
func PrintSobrenome( sobrenome string) string {
	return sobrenome
}