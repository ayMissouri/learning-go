package main

import (
	"fmt"
	"slices"
)

func main() {
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, len: %d, cap: %d\n", original, len(original), cap(original))

	s1 := original[2:5] // low:high (high is excluded)
	fmt.Printf("s1 (original[2:5]): %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	s2 := original[:4] // low:high (starting from 0)
	fmt.Printf("s2 (original[:4]): %v, len: %d, cap: %d\n", s2, len(s2), cap(s2))
	s3 := original[6:] // low:high (starting from 6 to end)
	fmt.Printf("s3 (original[6:]): %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	s4 := original[:] // full slice
	fmt.Printf("s4 (original[:]): %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))

	// does s4 contain 8?
	slices.Contains(s4, 8)
	fmt.Printf("s4: %v, len: %d, cap: %d\n", s4, len(s4), cap(s4))
}
