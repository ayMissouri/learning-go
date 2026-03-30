package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func config(numbers ...int) {
	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println(first)
	} else {
		fmt.Println("Default number")
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	// returns default value
	config()
}
