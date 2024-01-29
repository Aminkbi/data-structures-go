package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.Length++
}

func (l *LinkedList) DeleteNode(val int) {
	if l.Length == 0 {
		return
	}
	if l.Head.Data == val {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	prev := l.Head
	for prev.Next.Data != val {
		prev = prev.Next
		if prev.Next == nil {
			fmt.Println("Value not found in the list")
			return
		}
	}

	prev.Next = prev.Next.Next
	l.Length--
}

func (l LinkedList) PrintLinkedList() {

	nodeToPrint := l.Head

	for l.Length != 0 {
		fmt.Printf("%d", nodeToPrint.Data)
		nodeToPrint = nodeToPrint.Next
		l.Length--
		fmt.Print(" ")
	}
	fmt.Println()

}
