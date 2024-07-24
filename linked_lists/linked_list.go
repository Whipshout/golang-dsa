package linked_lists

import "errors"

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) PushFront(value T) {
	newNode := &Node[T]{value: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head.prev = newNode
		ll.head = newNode
	}

	ll.size++
}

func (ll *LinkedList[T]) PushBack(value T) {
	newNode := &Node[T]{value: value}

	if ll.tail == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}

	ll.size++
}

func (ll *LinkedList[T]) PopFront() (T, error) {
	if ll.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}

	value := ll.head.value
	oldHead := ll.head
	ll.head = ll.head.next

	if ll.head == nil {
		ll.tail = nil
	} else {
		ll.head.prev = nil
	}

	oldHead.next = nil
	ll.size--

	return value, nil
}

func (ll *LinkedList[T]) PopBack() (T, error) {
	if ll.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}

	value := ll.tail.value
	oldTail := ll.tail
	ll.tail = ll.tail.prev

	if ll.tail == nil {
		ll.head = nil
	} else {
		ll.tail.next = nil
	}

	ll.size--
	oldTail.prev = nil

	return value, nil
}

func (ll *LinkedList[T]) Front() (T, error) {
	if ll.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}

	return ll.head.value, nil
}

func (ll *LinkedList[T]) Back() (T, error) {
	if ll.IsEmpty() {
		var zero T
		return zero, errors.New("empty list")
	}

	return ll.tail.value, nil
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.size = 0
}
