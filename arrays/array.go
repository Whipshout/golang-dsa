package arrays

import "fmt"

type Array[T any] struct {
	data     []T
	size     int
	capacity int
}

func New[T any](initialCapacity ...int) *Array[T] {
	capacity := 10

	if len(initialCapacity) > 0 && initialCapacity[0] > 0 {
		capacity = initialCapacity[0]
	}

	return &Array[T]{
		data:     make([]T, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (a *Array[T]) Append(element T) {
	if a.size == a.capacity {
		a.resize()
	}

	a.data = append(a.data, element)
	a.size++
}

func (a *Array[T]) Get(index int) (T, error) {
	if index < 0 || index >= a.size {
		var zero T
		return zero, fmt.Errorf("index out of range")
	}

	return a.data[index], nil
}

func (a *Array[T]) Set(index int, element T) error {
	if index < 0 || index >= a.size {
		return fmt.Errorf("index out of range")
	}

	a.data[index] = element

	return nil
}

func (a *Array[T]) Size() int {
	return a.size
}

func (a *Array[T]) Pop() (T, error) {
	if a.size == 0 {
		var zero T
		return zero, fmt.Errorf("cannot pop from an empty array")
	}

	a.size--
	element := a.data[a.size]
	a.data = a.data[:a.size]

	if a.size < a.capacity/4 && a.capacity%4 > 10 {
		a.resize()
	}

	return element, nil
}

func (a *Array[T]) resize() {
	newCapacity := a.capacity * 2
	newData := make([]T, a.size, newCapacity)
	copy(newData, a.data)

	a.data = newData
	a.capacity = newCapacity
}
