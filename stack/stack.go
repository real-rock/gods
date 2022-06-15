package stack

import "errors"

type node struct {
	val      interface{}
	nextNode *node
	prevNode *node
}

func newNode(val interface{}) *node {
	return &node{
		val:      val,
		nextNode: nil,
		prevNode: nil,
	}
}

type Stack struct {
	top    *node
	length int
}

func New() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
	}
}

func (s *Stack) Push(val interface{}) {
	n := newNode(val)
	if s.top == nil {
		s.top = n
	} else {
		s.top.nextNode = n
		n.prevNode = s.top
		s.top = n
	}
	s.length++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("empty stack error")
	}
	n := s.top.val
	s.top = s.top.prevNode
	s.length--
	return n, nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Len() int {
	return s.length
}
