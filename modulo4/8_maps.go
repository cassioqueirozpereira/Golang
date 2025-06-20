// 2 - Maps: Heterogêneos
// Pode misturar tipos
// Estrutura chave - valor
// [key] = value
// Chave tem um tipo e o valor pode outro
// map[k]v -> k = chave, v = valor

// map[string]int
// {"Cássio": 33, "Dilnara": 30}
// map[string]sgrint
// {"Cássio": "Pereira", "Dilnara": "Pereira"}

package main

import "fmt"

func main() {
	fmt.Println("\nCriando um map e atribuindo valor depois\n")
	idade := map[string]int{}
	idade["Cássio"] = 33
	idade["Dilnara"] = 30
	
	fmt.Println(idade)
	fmt.Println(idade["Cássio"])
	fmt.Println(idade["Dilnara"])

	fmt.Println("\nCriando um map já com valores atribuidos\n")
	anoNasc := map[string]int {
		"Cássio": 1992,
		"Dilnara": 1995,
	}
	
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Cássio"])
	fmt.Println(anoNasc["Dilnara"])
	anoNasc["Victor"] = 2010
	fmt.Println(anoNasc)
}