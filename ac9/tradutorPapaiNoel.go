package main

import (
	"bufio"
	"fmt"
	"os"
)

var paises = []string{"brasil", "alemanha", "austria", "coreia", "espanha", "grecia", "estados-unidos",
	"inglaterra", "australia", "portugal", "suecia", "turquia", "argentina", "chile",
	"mexico", "antardida", "canada", "irlanda", "belgica", "italia", "libia", "siria",
	"marrocos", "japao"}

var frase = []string{"Feliz Natal!", "Frohliche Weihnachten!", "Frohe Weihnacht!", "Chuk Sung Tan!",
	"Feliz Navidad!", "Kala Christougena!", "Merry Christmas!", "Merry Christmas!",
	"Merry Christmas!", "Feliz Natal!", "God Jul!", "Mutlu Noeller", "Feliz Navidad!",
	"Feliz Navidad!", "Feliz Navidad!", "Merry Christmas!", "Merry Christmas!", "Nollaig Shona Dhuit!",
	"Zalig Kerstfeest!", "Buon Natale!", "Buon Natale!", "Milad Mubarak!", "Milad Mubarak!", "Merii Kurisumasu!"}

type tradutor struct {
	pais     string
	saudacao string
}

var traducao []tradutor

func traduzir() {
	for i := range paises {
		traducao = append(traducao, tradutor{paises[i], frase[i]})
	}
}

func acharPais(pais string) (bool, int) {
	for i, p := range traducao {
		if p.pais == pais {
			return true, i
		}
	}
	return false, -1
}

func main() {
	traduzir()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		pais := scanner.Text()
		if found, pos := acharPais(pais); found {
			fmt.Println(traducao[pos].saudacao)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
	}
}
