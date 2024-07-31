package hash_tables

import (
	"fmt"
	"hash/fnv"
)

const initialSize = 16

type Node[K comparable, V any] struct {
	key   K
	value V
	next  *Node[K, V]
}

type HashTable[K comparable, V any] struct {
	buckets []*Node[K, V]
	size    int
}

func New[K comparable, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		buckets: make([]*Node[K, V], initialSize),
	}
}

func hash[K comparable](key K) int {
	h := fnv.New32a()
	switch v := any(key).(type) {
	case string:
		_, err := h.Write([]byte(v))
		if err != nil {
			return 0
		}
	case int:
		_, err := h.Write([]byte(fmt.Sprintf("%d", v)))
		if err != nil {
			return 0
		}
	default:
		_, err := h.Write([]byte(fmt.Sprintf("%v", v)))
		if err != nil {
			return 0
		}
	}

	return int(h.Sum32())
}

func (ht *HashTable[K, V]) Insert(key K, value V) {
	index := hash(key) % len(ht.buckets)
	node := ht.buckets[index]

	if node == nil {
		ht.buckets[index] = &Node[K, V]{key: key, value: value}
		ht.size++
		return
	}

	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		if node.next == nil {
			node.next = &Node[K, V]{key: key, value: value}
			ht.size++
			return
		}

		node = node.next
	}
}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	index := hash(key) % len(ht.buckets)
	node := ht.buckets[index]

	for node != nil {
		if node.key == key {
			return node.value, true
		}
		node = node.next
	}

	var zero V
	return zero, false
}

func (ht *HashTable[K, V]) Delete(key K) bool {
	index := hash(key) % len(ht.buckets)
	node := ht.buckets[index]

	if node == nil {
		return false
	}

	if node.key == key {
		ht.buckets[index] = node.next
		ht.size--
		return true
	}

	prev := node
	for node != nil {
		if node.key == key {
			prev.next = node.next
			ht.size--
			return true
		}
		prev = node
		node = node.next
	}

	return false
}

func (ht *HashTable[K, V]) Size() int {
	return ht.size
}

func (ht *HashTable[K, V]) Display() {
	for i, bucket := range ht.buckets {
		fmt.Printf("Bucket %d: ", i)
		node := bucket
		for node != nil {
			fmt.Printf("%v:%v -> ", node.key, node.value)
			node = node.next
		}
		fmt.Println("nil")
	}
}
