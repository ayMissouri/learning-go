package main

import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("start")
	// defer is executed after the function returns
	defer fmt.Println("deferred")
	fmt.Println("middle")
}

func lastInFirstOut() {
	fmt.Println("start")
	// defer is executed in LIFO order, so deferred 2 is executed first
	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")
	fmt.Println("middle")
}

func main() {
	defer func() {
		fmt.Println("before the return of main()")
	}()

	simpleDefer()
	fmt.Println("==================")
	lastInFirstOut()

	fmt.Println("Last line of main()")

	// example of defer with file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// make sure the file is closed
	defer file.Close()
}
