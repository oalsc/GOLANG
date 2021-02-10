package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Ola Mundo!"
	canal <- "Ola Mundo 2!"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
