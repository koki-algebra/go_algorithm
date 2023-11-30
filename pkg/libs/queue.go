package libs

import "container/list"

type Queue struct {
	l *list.List
}

func NewQueue() *Queue {
	return &Queue{
		l: list.New(),
	}
}

func (q *Queue) Enqueue(v any) {
	q.l.PushBack(v)
}

func (q *Queue) Dequeue() any {
	return q.l.Remove(q.l.Front())
}

func (q *Queue) IsEmpty() bool {
	return q.l.Len() == 0
}
