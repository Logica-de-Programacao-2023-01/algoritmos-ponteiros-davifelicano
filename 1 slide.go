// Escreva uma função swap que receba dois ponteiros para int como argumentos e troque
// os valores apontados por eles.
package main

import "fmt"

func swap(prt *int, prt2 *int) {
	v10 := *prt
	*prt = *prt2
	*prt2 = v10

}
func main() {
	var x int = 20
	var y int = 10

	fmt.Println("Antes da troca")
	fmt.Println(x)
	fmt.Println(y)
	swap(&x, &y)
	fmt.Println("Depois da troca:")
	fmt.Println(x)
	fmt.Println(y)
}
