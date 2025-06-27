package main

import "fmt"

func soma (x int) int {
	somas := 0
	for i := 0; i <= x; i++ {
		somas += i
	}
	return somas
}

func main() {
	fmt.Println(soma(10))
}