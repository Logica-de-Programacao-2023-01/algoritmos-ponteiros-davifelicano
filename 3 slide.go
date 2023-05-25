// Escreva uma função que receba um ponteiro para uma struct que contém informações
// de um produto (nome, preço e quantidade em estoque). A função deve atualizar o preço
// desse produto para um novo valor fornecido como argumento.
package main

import "fmt"

type produto struct {
	nome       string
	preço      float64
	quantidade int
}

func novop(p *produto, novopreço float64) {
	p.preço = novopreço
}
func main() {

	var novopreço float64 = 3.5
	troca := produto{
		nome:       "Biscoito",
		preço:      2.5,
		quantidade: 10,
	}
	novop(&troca, novopreço)
	fmt.Print(troca)
}
