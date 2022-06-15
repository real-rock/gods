package pqueue

import (
	"gods/commons"
)

const maxItem = 100

type Heap[T commons.Comparable] struct {
	vals    []T
	length  int
	maxItem int
}

func newHeap[T commons.Comparable](n int) *Heap[T] {
	if n < 0 {
		n = maxItem
	}
	return &Heap[T]{
		vals:    make([]T, n),
		length:  0,
		maxItem: n,
	}
}

func reheapDown[T commons.Comparable](values []T, root, bottom int) {
	var maxChild int
	leftChild := root*2 + 1
	rightChild := root*2 + 2
	if leftChild > bottom {
		return
	}
	if leftChild == bottom {
		maxChild = leftChild
	} else {
		if values[leftChild] > values[rightChild] {
			maxChild = leftChild
		} else {
			maxChild = rightChild
		}
	}
	if values[root] < values[maxChild] {
		commons.Swap(&values[root], &values[maxChild])
		reheapDown(values, maxChild, bottom)
	}
}

func (h *Heap[T]) ReheapDown() {
	reheapDown(h.vals, 0, h.length-1)
}

func reheapUp[T commons.Comparable](values []T, root, bottom int) {
	if bottom > root {
		parent := (bottom - 1) / 2
		if values[parent] < values[bottom] {
			commons.Swap(&values[parent], &values[bottom])
			reheapUp(values, root, parent)
		}
	}
}

func (h *Heap[T]) ReheapUp() {
	reheapUp(h.vals, 0, h.length-1)
}

func (h *Heap[T]) IsEmpty() bool {
	return h.length == 0
}

func (h *Heap[T]) IsFull() bool {
	return h.length == h.maxItem
}

func (h *Heap[T]) Len() int {
	return h.length
}
