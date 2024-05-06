package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	inicio *Node
	fim    *Node
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{value, nil}

	if q.fim == nil {
		q.inicio = newNode
		q.fim = newNode
	} else {
		q.fim.next = newNode
		q.fim = newNode
	}

	fmt.Println("Elemento inserido:", value)
}

func (q *Queue) Dequeue() int {
	if q.inicio == nil {
		fmt.Println("A fila está vazia")
		return -1
	}

	value := q.inicio.value
	q.inicio = q.inicio.next

	if q.inicio == nil {
		q.fim = nil
	}

	return value
}

func (q *Queue) Traverse() {
	if q.inicio == nil {
		fmt.Println("A fila está vazia")
		return
	}

	current := q.inicio
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	queue := Queue{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Elementos na fila:")
	queue.Traverse()

	fmt.Println("Elemento removido:", queue.Dequeue())

	fmt.Println("Elementos na fila:")
	queue.Traverse()

	queue.Enqueue(4)

	fmt.Println("Elementos na fila:")
	queue.Traverse()

	fmt.Println("Elemento removido:", queue.Dequeue())

	fmt.Println("Elementos na fila:")
	queue.Traverse()
}
