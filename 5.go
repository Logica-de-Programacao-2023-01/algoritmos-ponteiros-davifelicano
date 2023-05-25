// Escreva uma função em Go que receba um ponteiro para uma variável float64
// e atualize o valor da variável para a média aritmética entre o seu valor atual e o valor da constante Pi.
package main

import (
	"fmt"
	"math"
)

func ari(prt *float64) {
	a := *prt
	media := a + math.Pi/2
	*prt = media
}
func main() {
	var media float64 = 22
	ari(&media)
	fmt.Print(media)
}
