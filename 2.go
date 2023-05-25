// Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar.
// A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.
package main

import "fmt"

func parimpar(prt *int) {
	if *prt%2 == 0 {
		*prt = 0
	} else {
		*prt = 1
	}
}
func main() {
	var n int = 10
	parimpar(&n)
	fmt.Print(n)
}
