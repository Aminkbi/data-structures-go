package stack

type Stack struct {
	items []int
}

func (s *Stack) Push(num int) {
	s.items = append(s.items, num)
}

func (s *Stack) Pop() int {

	lastIndex := len(s.items) - 1

	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}
