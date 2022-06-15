package pqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	a := assert.New(t)

	pq := New[int](100)
	pq.Push(4)
	pq.Push(7)
	pq.Push(3)
	pq.Push(6)
	pq.Push(13)
	pq.Push(1)
	pq.Push(2)
	pq.Push(2)
	pq.Push(9)

	a.Equal(13, pq.vals[0], "root must be largest node")
	data := pq.Pop()
	a.Equal(13, data, "Pop() must be 13")

	a.Equal(9, pq.vals[0], "root must be largest node")
	data = pq.Pop()
	a.Equal(9, data, "Pop() must be 9")
}
