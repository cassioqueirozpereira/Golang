package main

import "fmt"

func main() {
	fmt.Println(indexParImpar("YCOLUE'VREER"))
}

func indexParImpar(palavra string) (string, string) {
	var palavraPar, palavraImpar string

	for i := 0; i < len(palavra); i++ {
		if i % 2 == 0 {
			palavraPar += string(palavra[i])
		} else {
			palavraImpar += string(palavra[i])
		}
	}

	return palavraPar, palavraImpar
}