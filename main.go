package main

import linked_list "data-structures-go/linked-list"

func main() {

	head := &linked_list.Node{
		Next: nil,
		Data: 3,
	}
	list := &linked_list.LinkedList{
		Head:   head,
		Length: 1,
	}

	newNode := &linked_list.Node{
		Data: 1,
		Next: nil,
	}

	list.Prepend(newNode)

	list.DeleteNode(11)
}
