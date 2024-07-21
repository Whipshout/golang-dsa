package main

import (
	"fmt"
	"golang-dsa/arrays"
)

func main() {
	// --------------------- Array ---------------------
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
	// --------------------- Array ---------------------
}
