package queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(num int) {
	q.items = append(q.items, num)
}

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
