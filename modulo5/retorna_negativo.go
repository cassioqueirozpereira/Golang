package main

import "fmt"

func retornaNegativo (x float64) float64 {
	if x <= 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	fmt.Println(retornaNegativo(5.678))
}