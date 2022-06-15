package pqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeap(t *testing.T) {
	a := assert.New(t)

	h := newHeap[int](100)
	h.vals[0] = 13
	h.vals[1] = 7
	h.vals[2] = 3
	h.vals[3] = 6
	h.vals[4] = 4
	h.vals[5] = 1
	h.vals[6] = 2
	h.vals[7] = 2
	h.vals[8] = 9

	h.length = 9

	h.ReheapUp()
	a.Equal(9, h.vals[1], "h.vals[1] must be 9")
	h.vals[0] = 11
	h.vals[2] = 13

	h.ReheapDown()
	a.Equal(13, h.vals[0], "root node must be 13")
	a.Equal(11, h.vals[2], "h.vals[2] node must be 11")
}
