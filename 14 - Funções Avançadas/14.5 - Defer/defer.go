package main

import "fmt"

func calculaNota(n1, n2 float32) bool {
	fmt.Println("Entrando no calculo de notas")
	defer fmt.Println("O resultado das notas foi")

	media := n1 + n2/2

	if media > 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(calculaNota(5, 8))
}
