// Escreva uma função em Go que receba um ponteiro para uma variável inteira
// e atualize o valor da variável para a soma dos valores dos seus dois últimos dígitos
// (por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).
package main

import "fmt"

func somault(prt *int) {
	u := *prt
	f1 := u % 10
	f2 := (u / 10) % 10
	*prt = f1 + f2
}
func main() {
	var numeroparasomar int = 1234
	somault(&numeroparasomar)
	fmt.Print(numeroparasomar)
}
