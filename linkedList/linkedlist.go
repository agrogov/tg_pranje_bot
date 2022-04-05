package linkedList

import (
	"fmt"
	it "tg_pranje_bot/iterator"
	"tg_pranje_bot/node"
)

type iterator struct {
	current *node.Node
}

func (it *iterator) Next() *node.Node {
	current := it.current
	it.current = it.current.Next()
	return current
}

func (it *iterator) HasNext() bool {
	if it.current == nil {
		return false
	}

	return true
}

type LinkedList struct {
	head    *node.Node
	current *node.Node
	length  int
}

func (li *LinkedList) Iterator() it.IIterator {
	head := li.head
	return &iterator{head}
}

func New(head, current *node.Node) *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func (li *LinkedList) Head() *node.Node {
	return li.head
}

// Prints elements from the list
func (li *LinkedList) Print() {
	head := li.head
	for head != nil {
		fmt.Printf("%v\n", head.Item())
		head = head.Next()
	}
}

// Inserts at front of the list
func (li *LinkedList) Insert(item string) *node.Node {
	newNode := node.New(item) // init node
	if li.head == nil {
		li.head = newNode
		li.current = newNode
		li.length++
		return newNode
	}

	newNode.SetNext(li.current)
	li.current = newNode
	li.head = newNode
	li.length = li.length + 1
	return newNode
}

// Inserts at the end of list
func (li *LinkedList) Append(item string) *node.Node {
	newNode := node.New(item) // new Node
	if li.head == nil {
		li.head = newNode
		li.current = newNode
		return newNode
	}

	li.current.SetNext(newNode)
	li.current = newNode
	li.length = li.length + 1
	return newNode
}

// Removes an element from front of the list
func (li *LinkedList) RemoveFront() *node.Node {
	head := li.head
	if head != nil {
		nextNode := head.Next()
		li.head = nextNode
		head = nil
	}
	return nil
}

// // Removes an element
// func (li *LinkedList) Remove(item int) *node.Node {

// }

// Removes an element from back of the list
// TODO: implement remove back from linked list
// func (li *LinkedList) RemoveBack() *node.Node {

// }
