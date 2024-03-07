package binary_search_tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(val int) {
	toBeAddedNode := &Node{
		Value: val,
		Left:  nil,
		Right: nil,
	}

	if val > n.Value {
		if n.Right == nil {
			n.Right = toBeAddedNode
			return
		}
		n.Right.Insert(val)
	} else {
		if n.Left == nil {
			n.Left = toBeAddedNode
			return
		}
		n.Left.Insert(val)
	}
}

func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value == n.Value {
		return true
	} else if value > n.Value {
		if n.Right == nil {
			return false
		}
		return n.Right.Search(value)
	} else if value < n.Value {
		if n.Left == nil {
			return false
		}
		return n.Left.Search(value)
	}
	return false
}
