package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrDivideByZero = errors.New("cannot divide by zero")
var ErrNumberTooLarge = errors.New("number too large")

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func newOpError(op string, code int, message string, t time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    t,
	}
}

func (op OpError) Error() string {
	return op.Message
}

func OpExample() error {
	return newOpError("doSomethingCode", 100, "something went wrong", time.Now())
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	if a > 1000 {
		return 0, ErrNumberTooLarge
	}

	return a / b, nil
}

func main() {
	value, err := divide(9999, 1)
	if err != nil {
		if errors.Is(err, ErrNumberTooLarge) {
			fmt.Println("error for number too large")
		} else if errors.Is(err, ErrDivideByZero) {
			fmt.Println("error for divide by zero")
		}
		return
	}
	fmt.Println(value)
}
