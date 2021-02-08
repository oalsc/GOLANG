package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("ola mundo", canal)

	fmt.Println("depois da funcao escrever")

	// Method 1
	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// method 2
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
