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
	q.l.PushFront(v)
}

func (q *Queue) Dequeue() any {
	e := q.l.Back()
	return q.l.Remove(e)
}

func (q *Queue) IsEmpty() bool {
	return q.l.Len() == 0
}
