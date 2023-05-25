// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
// A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.
package main

import "fmt"

func soma(prt *int, n int) {
	s := 0
	for i := 0; i < n; i++ {
		s += i
	}
	*prt = s

}
func main() {
	var n int = 5
	var r int
	soma(&r, n)
	fmt.Print(r)
}
