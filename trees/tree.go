package trees

import "fmt"

type Node[K comparable, V any] struct {
	key      K
	value    V
	children []*Node[K, V]
}

type Tree[K comparable, V any] struct {
	Root *Node[K, V]
}

func New[K comparable, V any](rootKey K, rootValue V) *Tree[K, V] {
	return &Tree[K, V]{
		Root: &Node[K, V]{key: rootKey, value: rootValue},
	}
}

func (n *Node[K, V]) AddChild(childKey K, childValue V) *Node[K, V] {
	child := &Node[K, V]{key: childKey, value: childValue}
	n.children = append(n.children, child)

	return child
}

func (n *Node[K, V]) Find(key K) *Node[K, V] {
	if n.key == key {
		return n
	}

	for _, child := range n.children {
		if found := child.Find(key); found != nil {
			return found
		}
	}

	return nil
}

func (n *Node[K, V]) RemoveChild(key K) bool {
	for i, child := range n.children {
		if child.key == key {
			n.children = append(n.children[:i], n.children[i+1:]...)
			return true
		}

		for _, subChild := range child.children {
			if subChild.RemoveChild(key) {
				return true
			}
		}
	}

	return false
}

func (n *Node[K, V]) Display(level int) {
	fmt.Printf("%s%v: %v\n", string(rune(' '+level*2)), n.key, n.value)

	for _, child := range n.children {
		child.Display(level + 1)
	}
}
