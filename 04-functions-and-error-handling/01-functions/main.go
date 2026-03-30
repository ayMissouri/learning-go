package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello %s\n", name)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

// returns the area of a rectangle as float64
func calculateArea(width, height float64) float64 {
	if width <= 0 || height <= 0 {
		fmt.Println("Invalid dimensions")
		return 0
	}
	return width * height
}

// function running itself
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	greet("John")
	add(10, 20)

	area := calculateArea(4, 4)
	fmt.Println(area)

	fmt.Println(factorial(5))

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//logger \:= func(message string) {
	//	fmt.Println(message)
	//}
	//
	//logger("Hello")
	//logger("World")
}
