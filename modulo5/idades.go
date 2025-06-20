package main

import "fmt"

// Minha função
func idades (humano int) [3]int {
	var idade [3]int
	for i := 0; i <= humano; i++ {
		if i == 0 {
			idade[0] += 0
			idade[1] += 0
			idade[2] += 0
		} else if i == 1 {
			idade[0] += 1
			idade[1] += 15
			idade[2] += 15
		} else if i == 2 {
			idade[0] += 1
			idade[1] += 9
			idade[2] += 9
			} else {
			idade[0] += 1
			idade[1] += 4
			idade[2] += 5
		}
	}
	return idade
}

// Melhor função
func calculateYears(years int) (result [3]int) {
	switch years {
	case 1: result = [3]int {1, 15, 15}
	case 2: result = [3]int {2, 24, 24}
	default: result = [3]int {years, 24 + 4 * (years - 2), 24 + 5 * (years -2)}
	}
	return
}

func main() {
	// Função que eu fiz
	fmt.Println(idades(10))
	// Melhor maneira de fazer
	fmt.Println(calculateYears(10))
}