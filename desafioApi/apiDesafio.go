package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Livro struct {
	Id     string `json:"id"`
	Isbn   string `json:"isbn"`
	Titulo string `json:"titulo"`
	Autor  *Autor `json:"autor"`
}

type Autor struct {
	Nome      string `json: "nome"`
	Sobrenome string `json: "sobrenome"`
}

var livros []Livro

func getLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(livros)
}

func criarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var livro Livro
	_ = json.NewDecoder(r.Body).Decode(&livro)
	livro.Id = strconv.Itoa(len(livros) + 1)
	livros = append(livros, livro)
	json.NewEncoder(w).Encode(livro)
}

func atualizarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range livros {
		if item.Id == params["id"] {
			livros = append(livros[:index], livros[index+1:]...)
			var livro Livro
			_ = json.NewDecoder(r.Body).Decode(&livro)
			livro.Id = params["id"]
			livros = append(livros, livro)
			json.NewEncoder(w).Encode(livro)
			return
		}
	}
}

func deletarLivro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	params := mux.Vars(r)
	for index, item := range livros {
		if item.Id == params["id"] {
			livros = append(livros[:index], livros[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(livros)
}

func main() {
	r := mux.NewRouter()

	livros = append(livros, Livro{Id: "1", Isbn: "448743", Titulo: "O Livro da Engenharia", Autor: &Autor{Nome: "João", Sobrenome: "Silva"}})
	livros = append(livros, Livro{Id: "2", Isbn: "853250", Titulo: "A Arte da Guerra", Autor: &Autor{Nome: "Sun", Sobrenome: "Tzu"}})

	r.HandleFunc("/livros", getLivros).Methods("GET")
	r.HandleFunc("/livros", criarLivro).Methods("POST")
	r.HandleFunc("/livros/{id}", atualizarLivro).Methods("PUT")
	r.HandleFunc("/livros/{id}", deletarLivro).Methods("DELETE")

	log.Println("Servidor iniciado na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
