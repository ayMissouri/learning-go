package main

import "fmt"

func main() {
	names := []string{"John", "Mark", "Bob"}
	fmt.Println(names)

	items := make([]int, 3, 5)

	fmt.Printf("Items: %v, Len: %d, Cap: %d\n", items, len(items), cap(items))
	items = append(items, 1)
	items = append(items, 2)
	items = append(items, 3)
	// Capacity doubles when the slice is full
	fmt.Printf("Items: %v, Len: %d, Cap: %d\n", items, len(items), cap(items))

	// Slicing
	fmt.Printf("Slice: %v\n", items[3:6])
}
