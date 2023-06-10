package stack

// SliceStack Stack based on slice
type SliceStack[E any] struct {
	es []E
}

func (s *SliceStack[E]) Push(e E) E {
	s.es = append(s.es, e)
	return e
}

// Pop Removes the object at the top of this stack and returns that object as the value of this function.
// 注: 方法无法只返回一个值E，因为泛型E对应的具体类型可能无法用nil表示，故返回两个值，第二个值bool表示Pop操作是否成功
func (s *SliceStack[E]) Pop() (e E, b bool) {
	l := s.Len()
	if l == 0 {
		return
	}
	i := l - 1
	e = s.es[i]
	s.es = s.es[:i]
	return e, true
}

func (s *SliceStack[E]) Peek() (e E, b bool) {
	l := s.Len()
	if l == 0 {
		return
	}
	e = s.es[l-1]
	return e, true
}

func (s *SliceStack[E]) Len() int {
	return len(s.es)
}

func (s *SliceStack[E]) IsEmpty() bool {
	return s.Len() == 0
}

func NewSliceStack[E any](size int) *SliceStack[E] {
	s := new(SliceStack[E])
	if size > 0 {
		s.es = make([]E, 0, size)
	}
	return s
}
