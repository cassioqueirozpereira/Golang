// Structs
// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipo diferentes

package main

import "fmt"

// type <nome da nossa estrutura> struct { <campos> }
type Pessoa struct {
	Nome string
	Idade int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	println("\nStruct - Printando e definindo valores para a struct\n")
	fmt.Println(Pessoa{"Cássio", 33})
	fmt.Println(Pessoa{Nome: "Cássio", Idade: 33})
	fmt.Println(Pessoa{Nome: "Dilnara"})
	
	println("\nStruct - Atrelando a struct a uma variável\n")
	p1 := Pessoa{Nome: "Lanna", Idade: 12}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)
	
	println("\nStruct - Mudando o valor da struct através da variável\n")
	p1.Idade = 13
	fmt.Println(p1.Idade)
	
	p2 := Pessoa{Nome: "Victor", Idade: 14}
	
	println("\nStruct - Slice de pessoas\n")
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)
	
	println("\nStruct + map\n")
	alunos := map[string][]Pessoa{}
	alunos["Programação"] = pessoas
	fmt.Println(alunos)
	
	alunos2 := map[string][]Pessoa{
		"Programação": {{Nome: "Cássio", Idade: 33}, {Nome: "Victor", Idade: 14}},
		"Financeiro": {{Nome: "Dilnara", Idade: 30}, {Nome: "Lanna", Idade: 12}},
	}
	fmt.Println(alunos2)
	
	println("\nStruct - Herdando campos de outra struct\n")
	profissao := Profissao{p2, "dev"}
	fmt.Println(profissao)
	fmt.Println(profissao.Pessoa.Nome)
	fmt.Println(profissao.Pessoa.Idade)
	fmt.Println(profissao.Tipo)
}