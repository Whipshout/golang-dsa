package main

import (
	"fmt"
	"golang-dsa/arrays"
	"golang-dsa/queues"
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
}
