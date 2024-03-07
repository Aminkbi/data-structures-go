package main

import (
	binarysearchtree "data-structures-go/binary-search-tree"
	"fmt"
)

func main() {

	bst := &binarysearchtree.Node{
		Value: 5,
		Left:  nil,
		Right: nil,
	}
	bst.Insert(4)
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(1)

	bst.Insert(10)
	fmt.Println(bst.Search(10))

}
