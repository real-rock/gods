package queue

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

type Queue struct {
	front  *node
	rear   *node
	length int
}

func New() *Queue {
	return &Queue{
		front:  nil,
		rear:   nil,
		length: 0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.rear == nil
}

func (q *Queue) Len() int {
	return q.length
}

func (q *Queue) Push(val interface{}) {
	n := newNode(val)
	if q.rear == nil {
		q.front = n
		q.rear = n
	} else {
		q.rear.nextNode = n
		n.prevNode = q.rear
		q.rear = n
	}
	q.length++
}

func (q *Queue) Pop() error {
	if q.IsEmpty() {
		return errors.New("empty queue")
	}
	q.front = q.front.nextNode
	if q.front == nil {
		q.rear = nil
	}
	q.length--
	return nil
}

func (q *Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	return q.front.val, nil
}
