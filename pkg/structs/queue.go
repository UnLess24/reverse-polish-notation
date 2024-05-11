package structs

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) PeekFirst() T {
	var item T
	if len(q.data) > 0 {
		item = q.data[0]
	}
	return item
}

func (q *Queue[T]) PopFirst() T {
	var item T
	if len(q.data) > 0 {
		item = q.data[0]
		q.data = q.data[1:]
	}
	return item
}

func (q *Queue[T]) PushFirst(item T) {
	q.data = append([]T{item}, q.data...)
}

func (q *Queue[T]) PeekLast() T {
	var item T
	if len(q.data) > 0 {
		item = q.data[len(q.data)-1]
	}
	return item
}

func (q *Queue[T]) PopLast() T {
	var item T
	if len(q.data) > 0 {
		item = q.data[len(q.data)-1]
		q.data = q.data[:len(q.data)-1]
	}
	return item
}

func (q *Queue[T]) PushLast(item T) {
	q.data = append(q.data, item)
}
