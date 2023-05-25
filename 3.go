// Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.
package main

import "fmt"

func inver(prt *string) {
	r := []rune(*prt)
	t := len(r)
	for i := 0; i < t/2; i++ {
		r[i], r[t-1-i] = r[t-1-i], r[i]
	}
	*prt = string(r)
}
func main() {
	var letras string = "Juliana"
	inver(&letras)
	fmt.Print(letras)
}
