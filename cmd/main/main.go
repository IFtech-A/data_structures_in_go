package main

import "../../internal/lib/linkedlist"

func main() {
	root := linkedlist.New(1, "hello")
	linkedlist.Add(root, linkedlist.New(2, "world"))
	linkedlist.Add(root, linkedlist.New(3, "good"))
	linkedlist.Add(root, linkedlist.New(4, "to see"))
	linkedlist.PrintList(root)
}
