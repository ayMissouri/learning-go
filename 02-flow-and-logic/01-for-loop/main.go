package main

import "fmt"

func main() {
	// normal loop
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	// loop in the style of while loop
	j := 3
	for j > 0 {
		fmt.Println(j)
		j--
	}

	// loop in the style of for each
	items := []string{"apple", "banana", "orange"}
	// _ can be index, item can be _ if not needed
	for _, item := range items {
		fmt.Println(item)
	}
}
