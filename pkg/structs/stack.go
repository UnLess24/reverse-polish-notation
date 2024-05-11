package structs

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Peek() T {
	var item T
	if len(s.data) == 0 {
		return item
	}

	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Pop() T {
	var item T
	if len(s.data) == 0 {
		return item
	}

	item = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return item
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}
