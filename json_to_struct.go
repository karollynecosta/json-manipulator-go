package main

import (
	"encoding/json"
	"log"
)

// Transformando JSON em struct
// criar estrutura do que sera recebido via json

type User struct {
	FullName string
	Age      int `json:"age"`
}

func json_to_struct() {
	strJson := `{"fullname":"zunta", "age": 1}`

	// pegando a estrutura e colocando em variavel
	var val User

	err := json.Unmarshal([]byte(strJson), &val)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(val.Age)
}
