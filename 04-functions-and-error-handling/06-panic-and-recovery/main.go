package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in mightPanic")
	}

	fmt.Println("did not panic")
}

func recoverableFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic: ", r)
		}
	}()

	mightPanic(true)
}

func main() {
	recoverableFunc()
}
