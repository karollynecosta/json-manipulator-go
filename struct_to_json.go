package main

import (
	"encoding/json"
	"log"
)

// Transformando JSON em struct
// criar estrutura do que sera recebido via json

type User2 struct {
	FullName string `json:"fullname,omitempty"` // ignora o campo se ele estiver vazio
	Age      int    `json:"age"`
}

func main() {
	user := User2{
		FullName: "barto",
		Age:      3,
	}

	valor_byte, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(valor_byte))
}
