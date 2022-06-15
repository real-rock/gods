package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST(t *testing.T) {
	a := assert.New(t)

	b := New[int]()
	b.Insert(5)
	b.Insert(3)
	b.Insert(10)
	b.Insert(1)
	b.Insert(7)
	b.Insert(6)
	b.Insert(8)
	b.Insert(2)

	found := b.Search(3)
	a.True(found, "3 must be found")
	found = b.Search(4)
	a.False(found, "4 must be not found")

	b.Delete(3)
	found = b.Search(3)
	a.False(found, "3 must be not found")

	b.Print(InOrder)
	b.Print(PostOrder)
	b.Print(PreOrder)
}
