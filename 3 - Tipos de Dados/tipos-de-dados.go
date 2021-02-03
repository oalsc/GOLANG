package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = -100000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12300.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12.99
	fmt.Println(numeroReal3)

	var str string = "alo 123"
	fmt.Println(str)

	char := 'b'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
