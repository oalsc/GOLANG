package main

import "fmt"

func calculosMatematicos(x, y int) (soma int, subtracao int) {
	soma = x + y
	subtracao = x - y
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 15)
	fmt.Println(soma, subtracao)
}
