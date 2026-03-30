package main

import "fmt"

func main() {
	var numbers [2]int
	fmt.Printf("%v\n", numbers)

	numbers[0] = 10
	numbers[1] = 20
	fmt.Printf("%v\n", numbers)

	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%v\n", primes)
	primes[3] = 11
	fmt.Printf("%v\n", primes)

	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d\n", primes[i])
	}

	var matrix [2][3]int
	fmt.Printf("%v\n", matrix)

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][2] = 3
	fmt.Printf("%v\n", matrix)
}
