package linkedlist

import "fmt"

// listNode ...
type listNode struct {
	id   int
	data interface{}
	prev *listNode
	next *listNode
}

type List struct {
	size int

	head *listNode
	tail *listNode
}

// New creates new List structure
func New() *List {
	return new(List)
}

// Append - appends a given data to the end of the list
func (l *List) Append(data interface{}) {
	newNode := &listNode{
		data: data,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head

	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.size++
}

// Pop - pops the last element in the list.
func (l *List) Pop() interface{} {
	node := l.tail
	if node != nil {
		l.deleteNode(node)
		return node.data
	} else {
		return nil
	}
}

// InsertStart - inserts a data to the start of the list
func (l *List) InsertStart(data interface{}) {
	newNode := &listNode{
		data: data,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = l.head
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

// Size - returns the count of elements in the list
func (l *List) Size() int { return l.size }

func (l *List) deleteNode(node *listNode) {
	if node == nil {
		return
	}
	next := node.next
	prev := node.prev

	if next != nil {
		next.prev = prev
	}
	if prev != nil {
		prev.next = next
	}
	if l.head == node {
		l.head = next
	}
	if l.tail == node {
		l.tail = prev
	}
	l.size--
}

// Delete - removes the list element which has the data. If data is not unique in the list the first occurance will be deleted.
func (l *List) Delete(data interface{}) {
	for node := l.head; node != nil; node = node.next {
		if node.data != data {
			continue
		}
		l.deleteNode(node)
		break
	}
}

// Print - prints the data of the list elements
func (l *List) Print() {
	for node := l.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

// ReversePrint prints the data of the list elements in reverse order
func (l *List) ReversePrint() {
	for node := l.tail; node != nil; node = node.prev {
		fmt.Println(node.data)
	}
}

// Head - returns the data residing in the start of the list
func (l *List) Head() interface{} {
	if l.head == nil {
		return nil
	}
	return l.head.data
}

// Tail - returns the data residing in the end of the list
func (l *List) Tail() interface{} {
	if l.tail == nil {
		return nil
	}
	return l.tail.data
}
