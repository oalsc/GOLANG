package main

import (
	"clia/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inicialização")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
