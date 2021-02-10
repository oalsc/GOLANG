package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida dos estados")
	fmt.Println(tipoEndereco)
}
