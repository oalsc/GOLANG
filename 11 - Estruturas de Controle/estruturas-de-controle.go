package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 55

	if numero > 15 {
		fmt.Println(numero, "É maior que 15")
	} else {
		fmt.Println(numero, "É menor ou igual que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println(outroNumero, "É maior que 0")
	} else if outroNumero < -10 {
		fmt.Println(outroNumero, "É menor que -10")
	} else {
		fmt.Println(outroNumero, "Está entre 0 e -10")
	}

}
