package main

import "fmt"

func stringRepeat (x int, palavra string) string {
	var word string
	for i := 0; i < x; i++ {
		word += palavra
	}
	return word
}

func main() {
	fmt.Println(stringRepeat(5, "Cássio"))
}