package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_PushAndPop(t *testing.T) {
	a := assert.New(t)
	q := New()

	q.Push(1)
	q.Push(7)
	q.Push(5)
	q.Push(4)
	q.Push(6)
	q.Push(2)

	res, err := q.Front()
	a.NoError(err, "error should not be generated")
	a.Equal(1, res, "Pop() should return 1")

	err = q.Pop()
	a.NoError(err, "error should not be generated")

	res, err = q.Front()
	a.NoError(err, "error should not be generated")
	a.Equal(7, res, "Pop() should return 7")

	err = q.Pop()
	a.NoError(err, "error should not be generated")

	res, err = q.Front()
	a.NoError(err, "error should not be generated")
	a.Equal(5, res, "Pop() should return 5")

	err = q.Pop()
	a.NoError(err, "error should not be generated")

	res, err = q.Front()
	a.NoError(err, "error should not be generated")
	a.Equal(4, res, "Pop() should return 4")

	err = q.Pop()
	a.NoError(err, "error should not be generated")

	res, err = q.Front()
	a.NoError(err, "error should not be generated")
	a.Equal(6, res, "Pop() should return 6")

	err = q.Pop()
	a.NoError(err, "error should not be generated")

	res, err = q.Front()
	a.NoError(err, "error should not be generated")
	a.Equal(2, res, "Pop() should return 2")

	err = q.Pop()
	a.NoError(err, "error should not be generated")

	res, err = q.Front()
	a.Error(err, "error should be generated")
	a.Nil(res, "result must be nil")
}

func TestQueue_Len(t *testing.T) {
	a := assert.New(t)

	q := New()
	a.True(q.IsEmpty(), "IsEmpty() should return true")
	q.Push(5)
	a.Equal(1, q.Len(), "Len() should return 1")
	q.Push(3)
	a.Equal(2, q.Len(), "Len() should return 2")
	q.Push(2)
	a.Equal(3, q.Len(), "Len() should return 3")
	q.Push(1)
	a.Equal(4, q.Len(), "Len() should return 4")

	a.False(q.IsEmpty(), "IsEmpty() should return false")

	_ = q.Pop()
	a.Equal(3, q.Len(), "Len() should return 3")
	_ = q.Pop()
	a.Equal(2, q.Len(), "Len() should return 2")
	_ = q.Pop()
	a.Equal(1, q.Len(), "Len() should return 1")
	_ = q.Pop()
	a.Equal(0, q.Len(), "Len() should return 0")
}
