package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_PushAndPop(t *testing.T) {
	a := assert.New(t)
	s := New()

	s.Push(1)
	s.Push(7)
	s.Push(5)
	s.Push(4)
	s.Push(6)
	s.Push(2)

	res, err := s.Pop()
	a.NoError(err, "error should not be generated")
	a.Equal(2, res, "Pop() should return 2")

	res, err = s.Pop()
	a.NoError(err, "error should not be generated")
	a.Equal(6, res, "Pop() should return 6")

	res, err = s.Pop()
	a.NoError(err, "error should not be generated")
	a.Equal(4, res, "Pop() should return 4")

	res, err = s.Pop()
	a.NoError(err, "error should not be generated")
	a.Equal(5, res, "Pop() should return 5")

	res, err = s.Pop()
	a.NoError(err, "error should not be generated")
	a.Equal(7, res, "Pop() should return 7")

	res, err = s.Pop()
	a.NoError(err, "error should not be generated")
	a.Equal(1, res, "Pop() should return 1")

	res, err = s.Pop()
	a.Error(err, "error should be generated")
	a.Nil(res, "result must be nil")
}

func TestStack_Len(t *testing.T) {
	a := assert.New(t)

	s := New()
	a.True(s.IsEmpty(), "IsEmpty() should return true")
	s.Push(5)
	a.Equal(1, s.Len(), "Len() should return 1")
	s.Push(3)
	a.Equal(2, s.Len(), "Len() should return 2")
	s.Push(2)
	a.Equal(3, s.Len(), "Len() should return 3")
	s.Push(1)
	a.Equal(4, s.Len(), "Len() should return 4")

	a.False(s.IsEmpty(), "IsEmpty() should return false")

	_, _ = s.Pop()
	a.Equal(3, s.Len(), "Len() should return 3")
	_, _ = s.Pop()
	a.Equal(2, s.Len(), "Len() should return 2")
	_, _ = s.Pop()
	a.Equal(1, s.Len(), "Len() should return 1")
	_, _ = s.Pop()
	a.Equal(0, s.Len(), "Len() should return 0")
}
