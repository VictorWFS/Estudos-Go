package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Usuarios struct {
	Usuarios []Usuario `json: "Usuarios"`
}

type Usuario struct {
	Nome   string `json: "Nome"`
	Tipo   string `json: "Tipo"`
	Idade  int    `json: "Idade"`
	Social Social `json: "Social"`
}

type Social struct {
	facebook  string `json: Facebook`
	instagram string `json: Instagram`
}

func main() {
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo aberto com sucesso!")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Usuarios

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Usuarios); i++ {
		fmt.Println("Usuario tipo: " + users.Usuarios[i].Tipo)
		fmt.Println("Usuario Idade: " + strconv.Itoa(users.Usuarios[i].Idade))
		fmt.Println("Usuario Nome: " + users.Usuarios[i].Nome)
	}
}
