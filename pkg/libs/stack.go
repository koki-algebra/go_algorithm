package libs

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Push(v any) {
	s.l.PushFront(v)
}

func (s *Stack) Pop() any {
	if s.l.Len() == 0 {
		panic("stack is empty")
	}
	e := s.l.Front()
	return s.l.Remove(e)
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}
