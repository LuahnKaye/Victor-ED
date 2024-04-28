package main

import (
	"fmt"
)

func quebrarCifra(cifra string, desloque int) string {
	decifrado := ""
	for _, char := range cifra {
		if char >= 'A' && char <= 'Z' {
			decifradoChar := ((char-'A')+26-rune(desloque))%26 + 'A'
			decifrado += string(decifradoChar)
		} else {
			decifrado += string(char)
		}
	}
	return decifrado
}

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var cifra string
		var desloque int
		fmt.Scanln(&cifra)
		fmt.Scanln(&desloque)

		decifrado := quebrarCifra(cifra, desloque)
		fmt.Println(decifrado)
	}
}