package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "pedro",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Douglas",
			"ultimo":   "Souza",
		},
		"curso": {
			"nome":   "Engenharia Mecânica",
			"Campus": "campus b",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["signo"]["nome"])
}
