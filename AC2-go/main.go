package main

import (
	"AC2-go/geometria"
	"fmt"
	"math"
	"math/rand"
)

type Ponto struct {
	x int
	y int
}

func (ponto Ponto) DistanciaOrigem() {
	valor := (math.Sqrt(float64(ponto.x*ponto.x + ponto.y*ponto.y)))
	fmt.Printf("%.2f", valor)
}

func geraVetor() {

	var vetor [10]int

	for i := range vetor {
		vetor[i] = rand.Intn(100)
	}

	fmt.Println("Vetor gerado:")
	for _, v := range vetor {
		fmt.Println(v)
	}
}

func inverterString(input string) {
	runas := []rune(input)
	n := len(runas)

	for i := 0; i < n/2; i++ {
		runas[i], runas[n-1-i] = runas[n-1-i], runas[i]
	}

	fmt.Println("String invertida:", string(runas))
}

func main() {

	// AC2.1
	geraVetor()

	// AC2.2
	var input string
	fmt.Println("Digite uma string:")
	fmt.Scanln(&input)

	inverterString(input)

	// AC2.3
	ponto := Ponto{x: 3, y: 4}
	ponto.DistanciaOrigem()

	// AC2.4
	var largura, altura float64

	fmt.Print("Digite a largura do retângulo: ")
	fmt.Scanln(&largura)

	fmt.Print("Digite a altura do retângulo: ")
	fmt.Scanln(&altura)

	retangulo := geometria.Retangulo{Largura: largura, Altura: altura}

	area := retangulo.Area()
	perimetro := retangulo.Perimetro()

	fmt.Println("Área do retângulo:", area)
	fmt.Println("Perímetro do retângulo:", perimetro)
}
