package heap

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("Cannot extract from empty array")
		return -1
	}
	extracted := h.array[0]

	length := len(h.array) - 1

	h.swap(0, length)
	h.array = h.array[:length]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}
