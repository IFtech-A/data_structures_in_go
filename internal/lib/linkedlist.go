package linkedlist

import "fmt"

// LinkedList ...
type LinkedList struct {
	id   int
	data string
	next *LinkedList
}

// New ...
func New(id int, data string) *LinkedList {
	return &LinkedList{
		id:   id,
		data: data,
		next: nil,
	}
}

// AddToList ...
func AddToList(list *LinkedList, new *LinkedList) {
	var tempList *LinkedList = nil
	tempList = list
	for tempList.next != nil {
		tempList = tempList.next
	}
	tempList.next = new
}

// PrintList ...
func PrintList(list *LinkedList) {
	var tempList *LinkedList = nil
	tempList = list
	for tempList != nil {
		fmt.Printf("id %v,  data %v\n", tempList.id, tempList.data)
		tempList = tempList.next
	}
}
