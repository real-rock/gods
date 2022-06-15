package bst

import (
	"fmt"
	"gods/commons"
)

const (
	InOrder   = 0
	PostOrder = 1
	PreOrder  = 2
)

type treeNode[T commons.Comparable] struct {
	val        T
	leftChild  *treeNode[T]
	rightChild *treeNode[T]
}

func newNode[T commons.Comparable](val T) *treeNode[T] {
	return &treeNode[T]{
		val:        val,
		leftChild:  nil,
		rightChild: nil,
	}
}

type BST[T commons.Comparable] struct {
	root   *treeNode[T]
	length int
}

func New[T commons.Comparable]() *BST[T] {
	return &BST[T]{
		root:   nil,
		length: 0,
	}
}

func (b *BST[T]) IsEmpty() bool {
	return b.root == nil
}

func (b *BST[T]) Len() int {
	return b.length
}

func searchItem[T commons.Comparable](root *treeNode[T], item T) bool {
	if root == nil {
		return false
	}
	if root.val == item {
		return true
	} else if root.val < item {
		return searchItem(root.rightChild, item)
	} else {
		return searchItem(root.leftChild, item)
	}
}

func (b *BST[T]) Search(item T) bool {
	return searchItem(b.root, item)
}

func insertItem[T commons.Comparable](root **treeNode[T], item T) {
	if *root == nil {
		*root = newNode(item)
	} else if (*root).val <= item {
		insertItem(&(*root).rightChild, item)
	} else {
		insertItem(&(*root).leftChild, item)
	}
}

func (b *BST[T]) Insert(item T) {
	insertItem(&b.root, item)
}

func getPredecessor[T commons.Comparable](root *treeNode[T]) T {
	iter := root
	for iter.rightChild != nil {
		iter = iter.rightChild
	}
	return iter.val
}

func deleteNode[T commons.Comparable](root **treeNode[T]) {
	if (*root).leftChild == nil {
		*root = (*root).rightChild
	} else if (*root).rightChild == nil {
		*root = (*root).leftChild
	} else {
		data := getPredecessor((*root).rightChild)
		(*root).val = data
		deleteItem(&(*root).leftChild, data)
	}
}

func deleteItem[T commons.Comparable](root **treeNode[T], item T) {
	if item < (*root).val {
		deleteItem(&(*root).leftChild, item)
	} else if item > (*root).val {
		deleteItem(&(*root).rightChild, item)
	} else {
		deleteNode(root)
	}
}

func (b *BST[T]) Delete(item T) {
	deleteItem(&b.root, item)
}

func inTraversal[T commons.Comparable](root *treeNode[T], f func(args ...interface{})) {
	if root == nil {
		return
	}
	inTraversal(root.leftChild, f)
	f(root.val)
	inTraversal(root.rightChild, f)
}

func postTraversal[T commons.Comparable](root *treeNode[T], f func(args ...interface{})) {
	if root == nil {
		return
	}
	postTraversal(root.leftChild, f)
	postTraversal(root.rightChild, f)
	f(root.val)
}

func preTraversal[T commons.Comparable](root *treeNode[T], f func(args ...interface{})) {
	if root == nil {
		return
	}
	f(root.val)
	preTraversal(root.leftChild, f)
	preTraversal(root.rightChild, f)
}

func (b *BST[T]) Print(orderType int) {
	switch orderType {
	case InOrder:
		inTraversal(b.root, func(args ...interface{}) {
			fmt.Printf("%v ", args)
		})
	case PostOrder:
		postTraversal(b.root, func(args ...interface{}) {
			fmt.Printf("%v ", args)
		})
	case PreOrder:
		preTraversal(b.root, func(args ...interface{}) {
			fmt.Printf("%v ", args)
		})
	}
	fmt.Println()
}
