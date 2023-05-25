// Escreva uma função em Go que receba um ponteiro para um struct
// Livro com campos título e autor,
// e altere o título do livro para "Desconhecido" se o autor for "Anônimo".
package main

import "fmt"

type livro struct {
	título string
	autor  string
}

func alterar(l *livro) {

	if l.autor == "Anônimo" {
		l.título = "Desconhecido"
	}
}
func main() {
	book := livro{
		título: "Rei leão",
		autor:  "Anônimo",
	}
	fmt.Println(book)
	alterar(&book)
	fmt.Println(book)
}
