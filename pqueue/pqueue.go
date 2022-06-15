package pqueue

import (
	"gods/commons"
	"log"
)

type PriorityQueue[T commons.Comparable] struct {
	*Heap[T]
}

func New[T commons.Comparable](maxItem int) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		newHeap[T](maxItem),
	}
}

func (pq *PriorityQueue[T]) Push(item T) {
	if pq.IsFull() {
		log.Panicln("priority queue is already full")
	}
	pq.vals[pq.length] = item
	pq.length++
	pq.ReheapUp()
}

func (pq *PriorityQueue[T]) Pop() T {
	if pq.IsEmpty() {
		log.Panicln("priority queue is empty")
	}
	data := pq.vals[0]
	pq.vals[0] = pq.vals[pq.length-1]
	pq.length--
	pq.ReheapDown()
	return data
}
