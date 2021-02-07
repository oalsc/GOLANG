package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"João", "David", "Lucas"}

	// for _, nome := range nomes {
	// 	fmt.Println(nome)
	// }

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome      string
		sobrenome string
	}

	// Não funciona

	// usuario2 := usuarioStruct{"Zé", "Jr"}

	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }
}
