package queue

import (
	"errors"

	"github.com/MaestroShifu/concurrent-queue-golang/queue/node"
)

// Backend para el manejo de las colas
type QueueBasic struct {
	// Punto de inicio y fin
	head *node.Node
	tail *node.Node

	// Indicadores de tamaño
	size    uint32
	maxSize uint32
}

func NewQueueBasic(maxSize uint32) *QueueBasic {
	return &QueueBasic{
		size:    0,
		head:    nil,
		tail:    nil,
		maxSize: maxSize,
	}
}

func (queue *QueueBasic) Put(data interface{}) error {
	if queue.size >= queue.maxSize {
		err := errors.New("la cola esta llena")
		return err
	}

	node := node.NewNode(data)
	if queue.size == 0 {
		queue.head = node
		queue.tail = node
		queue.size++
		return nil
	}

	currentHead := queue.head
	newHead := node
	newHead.SetNext(currentHead)
	currentHead.SetPrev(newHead)
	queue.head = newHead
	queue.size++
	return nil
}

func (queue *QueueBasic) Pop() (interface{}, error) {
	if queue.size == 0 {
		err := errors.New("la cola esta vacia")
		return nil, err
	}

	currentEnd := queue.tail
	newEnd := currentEnd.GetPrev()

	if newEnd != nil {
		newEnd.SetNext(nil)
		queue.tail = newEnd
	}

	queue.size--
	if queue.size == 0 {
		queue.head = nil
		queue.tail = nil
	}

	return currentEnd.GetData(), nil
}

func (queue *QueueBasic) IsEmpty() bool {
	return queue.size == 0
}

func (queue *QueueBasic) IsFull() bool {
	return queue.size >= queue.maxSize
}
