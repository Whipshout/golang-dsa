package binary_trees

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T constraints.Ordered] struct {
	Root *Node[T]
}

func New[T constraints.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (bt *BinaryTree[T]) Insert(value T) {
	bt.Root = insert(bt.Root, value)
}

func insert[T constraints.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{Value: value}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}

	return node
}

func (bt *BinaryTree[T]) Search(value T) bool {
	return search(bt.Root, value)
}

func search[T constraints.Ordered](node *Node[T], value T) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return search(node.Left, value)
	}

	return search(node.Right, value)
}

func (bt *BinaryTree[T]) Delete(value T) {
	bt.Root = deleteNode(bt.Root, value)
}

func deleteNode[T constraints.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right = deleteNode(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minNode := findMin(node.Right)
		node.Value = minNode.Value
		node.Right = deleteNode(node.Right, minNode.Value)
	}

	return node
}

func findMin[T constraints.Ordered](node *Node[T]) *Node[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}

	return current
}

func (bt *BinaryTree[T]) InorderTraversal() []T {
	result := make([]T, 0)
	inorderTraversal(bt.Root, &result)

	return result
}

func inorderTraversal[T constraints.Ordered](node *Node[T], result *[]T) {
	if node != nil {
		inorderTraversal(node.Left, result)
		*result = append(*result, node.Value)
		inorderTraversal(node.Right, result)
	}
}
