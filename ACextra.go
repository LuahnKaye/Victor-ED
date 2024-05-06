package main

import "fmt"

type Nó struct {
    valor int
    próximo  *Nó
}

type Fila struct {
    início *Nó
    fim    *Nó
}

func (f *Fila) Enfileirar(valor int) {
    novoNó := &Nó{valor, nil}

    if f.fim == nil {
        f.início = novoNó
        f.fim = novoNó
    } else {
        f.fim.próximo = novoNó
        f.fim = novoNó
    }

    fmt.Println("Elemento inserido:", valor)
}

func (f *Fila) Desenfileirar() int {
    if f.início == nil {
        fmt.Println("A fila está vazia")
        return -1
    }

    valor := f.início.valor
    f.início = f.início.próximo

    if f.início == nil {
        f.fim = nil
    }

    return valor
}

func (f *Fila) Percorrer() {
    if f.início == nil {
        fmt.Println("A fila está vazia")
        return
    }

    atual := f.início
    for atual != nil {
        fmt.Println(atual.valor)
        atual = atual.próximo
    }
}

func main() {
    fila := Fila{}

    fila.Enfileirar(1)
    fila.Enfileirar(2)
    fila.Enfileirar(3)

    fmt.Println("Elementos na fila:")
    fila.Percorrer()

    fmt.Println("Elemento removido:", fila.Desenfileirar())

    fmt.Println("Elementos na fila:")
    fila.Percorrer()

    fila.Enfileirar(4)

    fmt.Println("Elementos na fila:")
    fila.Percorrer()

    fmt.Println("Elemento removido:", fila.Desenfileirar())

    fmt.Println("Elementos na fila:")
    fila.Percorrer()
}
