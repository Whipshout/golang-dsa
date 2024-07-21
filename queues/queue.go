package queues

import "errors"

type Queue[T any] struct {
	elements []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{elements: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	item := q.elements[0]
	q.elements = q.elements[1:]

	return item, nil
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}

	return q.elements[0], nil
}
