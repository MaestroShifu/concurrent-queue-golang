package queue

import "errors"

// Almacenamiento de los nodos de la cola
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// Backend para el manejo de las colas
type QueueBackend struct {
	// Punto de inicio y fin
	head *Node
	tail *Node

	// Indicadores de tamaño
	size    uint32
	maxSize uint32
}

func (queue *QueueBackend) createNode(data interface{}) *Node {
	node := Node{
		data: data,
		next: nil,
		prev: nil,
	}
	return &node
}

func (queue *QueueBackend) put(data interface{}) error {
	if queue.size >= queue.maxSize {
		err := errors.New("La cola esta llena")
		return err
	}

	if queue.size == 0 {
		node := queue.createNode(data)
		queue.head = node
		queue.tail = node
		queue.size++
		return nil
	}

	currentHead := queue.head
	newHead := queue.createNode(data)
	newHead.next = currentHead
	currentHead.prev = newHead

	queue.head = currentHead
	return nil
}

func (queue *QueueBackend) pop() (interface{}, error) {
	if queue.size == 0 {
		err := errors.New("La cola esta vacia")
		return nil, err
	}

	currentEnd := queue.tail
	newEnd := currentEnd.prev

	if newEnd != nil {
		newEnd.next = nil
	}

	queue.size--
	if queue.size == 0 {
		queue.head = nil
		queue.tail = nil
	}

	return currentEnd.data, nil
}

func (queue *QueueBackend) isEmpty() bool {
	return queue.size == 0
}

func (queue *QueueBackend) isFull() bool {
	return queue.size >= queue.maxSize
}
