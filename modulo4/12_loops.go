package main

import (
	"fmt"
	"time"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		sum += 1
	}
	
	frutas := []string {"laranja", "maçã", "banana", "uva", "kiwi"}
	
	for i, fruta := range frutas {
		fmt.Println("\nFruta:",fruta, "Índice:", i)
		time.Sleep(2 * time.Second)
	}
}