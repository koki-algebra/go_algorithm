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
	s.l.PushBack(v)
}

func (s *Stack) Pop() any {
	if s.l.Len() == 0 {
		panic("stack is empty")
	}
	return s.l.Remove(s.l.Back())
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}
