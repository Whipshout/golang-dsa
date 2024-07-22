package stacks

import "errors"

type Stack[T any] struct {
	elements []T
	capacity int
}

func New[T any](initialCapacity ...int) *Stack[T] {
	capacity := 10

	if len(initialCapacity) > 0 && initialCapacity[0] > 0 {
		capacity = initialCapacity[0]
	}

	return &Stack[T]{
		elements: make([]T, 0, capacity),
		capacity: capacity,
	}
}

func (s *Stack[T]) Push(element T) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.elements = append(s.elements, element)

	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]

	return element, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) IsFull() bool {
	return len(s.elements) == s.capacity
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Clear() {
	s.elements = s.elements[:0]
}
