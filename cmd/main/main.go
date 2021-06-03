package main

import (
	"fmt"

	"github.com/iftech-a/data_structures_in_go/internal/lib/linkedlist"
)

func main() {
	root := linkedlist.New()
	root.Append("hello")
	fmt.Println(root.Pop())
	fmt.Println(root.Pop())
	root.Print()
	root.ReversePrint()

}
