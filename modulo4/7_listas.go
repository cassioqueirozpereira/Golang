// Listas

// 1 - Arrays e Slices: Homogêneos
// Todos os elementos tem o mesmo tipo
// [1, 2, 3, 4, 5] - [int]
// ["Cássio", "Victor", "Lanna", "Dilnara"] - [string]

// 2 - Maps: Heterogêneos
// Pode misturar valor
// Estrutura chave - valor
// [key] = value
// Chave tem um tipo e o valor pode ter outro tipo
// map[string]int
// {"Cássio": 33, "Dilnara": 30}
// map[string]string
// {"Cássio": "Pereira", "Dinara": "Pereira"}

// Array

// Tamanho fixo, de zero ou mais elementos do mesmo tipo
// Acessamos os valores com índice: a[0], a[1]...
// Função embutida len(), retorna o tamanho do Array
// Por conta do tamanho fixo, não é muito utilizada. Só em casos específicos.

// Slice

// Parecido com o array, mas sem tamanho fixo
// Função append(), usada para adicionar valores

package main

import "fmt"

func main() {
	// Array - tamanho fixo
	println("Array - tamanho fixo\n")

	var array [2]string

	array[0] = "Hello"
	array[1] = "World!"

	println("\nArray - chamar o conteúdo separado e a lista toda\n")

	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[0], array[1])
	fmt.Println(array)
	
	println("\nArray - chamar conforme as posições\n")
	
	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:3])
	fmt.Println(numPrimos[:3])
	fmt.Println(numPrimos[3:])
	
	// Slices
	println("\nSlices - sem tamanho fixo\n")
	
	// Jeito errado
	// var slice []string
	slice := make([]string, 2)
	slice[0] = "Hello"
	slice[1] = "World!"
	
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	// fmt.Println(slice[2]) //Aqui vai dar erro, pois não temos uma terceira posição

	println("\nSlices - adicionando valores ao slice\n")

	numPares := []int {2, 4, 6}
	fmt.Println(numPares)

	numPares = append(numPares, 8, 10)
	fmt.Println(numPares)
}