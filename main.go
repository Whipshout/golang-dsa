package main

import (
	"fmt"
	"golang-dsa/arrays"
	"golang-dsa/hash_tables"
	"golang-dsa/heaps"
	"golang-dsa/linked_lists"
	"golang-dsa/queues"
	"golang-dsa/sets"
	"golang-dsa/stacks"
)

func main() {
	// --------------------- Array ---------------------
	fmt.Println("--------------------- Array ---------------------")
	intArray := arrays.New[int](10)
	intArray.Append(2)
	intArray.Append(4)
	fmt.Printf("intArray length: %v\n", intArray.Size())

	item, _ := intArray.Get(0)
	fmt.Printf("intArray[0]: %v\n", item)

	_, err := intArray.Get(2)
	fmt.Printf("intArray[2]: %v\n", err)

	_ = intArray.Set(0, 10)
	item, _ = intArray.Get(0)
	fmt.Printf("intArray[0] updated: %v\n", item)

	_ = intArray.Set(2, 10)
	_, err = intArray.Get(2)
	fmt.Printf("intArray[2]: %v\n", err)

	lastItem, _ := intArray.Pop()
	fmt.Printf("intArray Pop: %v\n", lastItem)
	fmt.Println("--------------------- Array ---------------------")
	// --------------------- Array ---------------------
	fmt.Println()
	// --------------------- Queue ---------------------
	fmt.Println("--------------------- Queue ---------------------")
	stringQueue := queues.New[string]()
	fmt.Printf("stringQueue is empty: %v\n", stringQueue.IsEmpty())
	stringQueue.Enqueue("Hello")
	stringQueue.Enqueue("World")
	fmt.Printf("stringQueue length: %v\n", stringQueue.Size())

	peakItem, _ := stringQueue.Peek()
	fmt.Printf("peakItem: %v\n", peakItem)

	stringItem, _ := stringQueue.Dequeue()
	fmt.Printf("First queue item: %v\n", stringItem)

	stringItem, _ = stringQueue.Dequeue()
	fmt.Printf("Second queue item: %v\n", stringItem)

	_, err = stringQueue.Peek()
	fmt.Printf("peakItem: %v\n", err)
	fmt.Println("--------------------- Queue ---------------------")
	// --------------------- Queue ---------------------
	fmt.Println()
	// --------------------- Stack ---------------------
	fmt.Println("--------------------- Stack ---------------------")
	booleanStack := stacks.New[bool]()
	fmt.Printf("booleanStack is empty: %v\n", booleanStack.IsEmpty())
	_ = booleanStack.Push(true)
	_ = booleanStack.Push(false)
	fmt.Printf("booleanStack length: %v\n", booleanStack.Size())

	peakStack, _ := booleanStack.Peek()
	fmt.Printf("peakStack: %v\n", peakStack)

	booleanItem, _ := booleanStack.Pop()
	fmt.Printf("Last stack item: %v\n", booleanItem)

	booleanItem, _ = booleanStack.Pop()
	fmt.Printf("Last stack item: %v\n", booleanItem)

	_, err = booleanStack.Peek()
	fmt.Printf("peakStack: %v\n", err)
	fmt.Println("--------------------- Stack ---------------------")
	// --------------------- Stack ---------------------
	fmt.Println()
	// --------------------- Linked List ---------------------
	fmt.Println("--------------------- Linked List ---------------------")
	intLL := linked_lists.New[int]()
	fmt.Printf("intLL is empty: %v\n", intLL.IsEmpty())
	intLL.PushFront(1)
	intLL.PushFront(2)
	intLL.PushFront(3)
	fmt.Printf("intLL length: %v\n", intLL.Size())

	frontValue, _ := intLL.Front()
	fmt.Printf("frontValue: %v\n", frontValue)

	backValue, _ := intLL.Back()
	fmt.Printf("backValue: %v\n", backValue)

	frontValue, _ = intLL.PopFront()
	fmt.Printf("frontValue: %v\n", frontValue)

	backValue, _ = intLL.PopBack()
	fmt.Printf("backValue: %v\n", backValue)

	frontValue, _ = intLL.PopFront()
	fmt.Printf("frontValue: %v\n", frontValue)
	fmt.Println("--------------------- Linked List ---------------------")
	// --------------------- Linked List ---------------------
	fmt.Println()
	// --------------------- Heap ---------------------
	fmt.Println("--------------------- Heap ---------------------")
	boolHeap := heaps.New[int]()
	fmt.Printf("boolHeap length: %v\n", boolHeap.Size())
	boolHeap.Push(1)
	boolHeap.Push(3)
	boolHeap.Push(2)
	fmt.Printf("boolHeap length: %v\n", boolHeap.Size())

	popValue, _ := boolHeap.Pop()
	fmt.Printf("popValue: %v\n", popValue)

	popValue, _ = boolHeap.Pop()
	fmt.Printf("popValue: %v\n", popValue)

	popValue, _ = boolHeap.Pop()
	fmt.Printf("popValue: %v\n", popValue)

	_, err = boolHeap.Pop()
	fmt.Printf("popValue: %v\n", err)
	fmt.Println("--------------------- Heap ---------------------")
	// --------------------- Heap ---------------------
	fmt.Println()
	// --------------------- Set ---------------------
	fmt.Println("--------------------- Set ---------------------")
	stringSet := sets.New[string]()
	fmt.Printf("stringSet length: %v\n", stringSet.Size())
	stringSet.Add("a")
	stringSet.Add("b")
	stringSet.Add("c")
	fmt.Printf("stringSet length: %v\n", stringSet.Size())

	stringSet2 := sets.New[string]()
	stringSet2.Add("b")
	stringSet2.Add("d")

	contains := stringSet2.Contains("a")
	fmt.Printf("contains: %v\n", contains)

	unionSet := stringSet.Union(*stringSet2)
	fmt.Printf("unionSet length: %v\n", unionSet.Size())

	intersectionSet := stringSet.Intersection(*stringSet2)
	fmt.Printf("intersectionSet length: %v\n", intersectionSet.Size())

	differenceSet := stringSet.Difference(*stringSet2)
	fmt.Printf("differenceSet length: %v\n", differenceSet)

	sliceSet := stringSet.ToSlice()
	fmt.Printf("sliceSet length: %v\n", len(sliceSet))

	stringSet.Remove("a")
	fmt.Printf("stringSet length: %v\n", stringSet.Size())

	stringSet.Clear()
	fmt.Printf("stringSet length: %v\n", stringSet.Size())
	fmt.Println("--------------------- Set ---------------------")
	// --------------------- Set ---------------------
	fmt.Println()
	// --------------------- Hash Table ---------------------
	fmt.Println("--------------------- Hash Table ---------------------")
	htInt := hash_tables.New[string, int]()
	htInt.Insert("one", 1)
	htInt.Insert("two", 2)
	fmt.Printf("htInt length: %v\n", htInt.Size())

	intValue, _ := htInt.Get("two")
	fmt.Printf("Key two value: %v\n", intValue)

	htInt.Display()

	htInt.Delete("two")
	_, found := htInt.Get("two")
	fmt.Printf("Key two found: %v\n", found)

	fmt.Println("--------------------- Hash Table ---------------------")
	// --------------------- Hash Table ---------------------
	fmt.Println()
}
