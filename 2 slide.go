// Escreva uma função que receba um ponteiro para um booleano e altere o valor desse
// booleano para o seu inverso.
package main

import "fmt"

func inv(prt *bool) {
	if *prt == true {
		*prt = false
	} else if *prt == false {
		*prt = true
	}
}
func main() {
	var x bool = false
	fmt.Println("antes de trocar:")
	fmt.Println(x)
	fmt.Println("Depois de trocar")
	inv(&x)
	fmt.Println(x)
}
