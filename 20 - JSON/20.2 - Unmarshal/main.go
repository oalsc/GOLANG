package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2JSON := `{"nome":"Toby","raca":"Poodle","idade":5}`
	c2 := make(map[string]interface{})
	if erro := json.Unmarshal([]byte(cachorro2JSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2["nome"])
}
