package main

import (
	"data-structures-go/heap"
	"fmt"
)

func main() {

	m := &heap.MaxHeap{}

	buildHeap := []int{10, 20, 30, 40, 50}

	for _, val := range buildHeap {
		m.Insert(val)
		fmt.Println(m)
	}

}
