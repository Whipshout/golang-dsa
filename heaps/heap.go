package heaps

import (
	"errors"
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	elements []T
}

func New[T constraints.Ordered]() *Heap[T] {
	return &Heap[T]{elements: make([]T, 0)}
}

func (h *Heap[T]) Push(element T) {
	h.elements = append(h.elements, element)
	h.upHeapify(h.Size() - 1)
}

func (h *Heap[T]) Pop() (T, error) {
	if h.Size() == 0 {
		var zero T
		return zero, errors.New("empty heap")
	}

	minValue := h.elements[0]
	lastIndex := h.Size() - 1

	if lastIndex > 0 {
		h.elements[0] = h.elements[lastIndex]
		h.elements = h.elements[:lastIndex]
		h.downHeapify(0)
	} else {
		h.elements = h.elements[:0]
	}

	return minValue, nil
}

func (h *Heap[T]) Size() int {
	return len(h.elements)
}

func (h *Heap[T]) upHeapify(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.elements[parentIndex] <= h.elements[index] {
			break
		}
		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
		index = parentIndex
	}
}

func (h *Heap[T]) Peek() (T, error) {
	if h.Size() == 0 {
		var zero T
		return zero, errors.New("empty heap")
	}
	return h.elements[0], nil
}

func (h *Heap[T]) downHeapify(index int) {
	for {
		smallest := index
		leftChild := 2*index + 1
		rightChild := 2*index + 2

		if leftChild < h.Size() && h.elements[leftChild] < h.elements[smallest] {
			smallest = leftChild
		}
		if rightChild < h.Size() && h.elements[rightChild] < h.elements[smallest] {
			smallest = rightChild
		}

		if smallest == index {
			break
		}

		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		index = smallest
	}
}
