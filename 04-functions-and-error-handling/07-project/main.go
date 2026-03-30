package main

import "fmt"

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const (
	division      = "Division"
	divisionError = "Cannot divide by zero"
)

func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Error performing %s: %s", e.Operation, e.Message)
}

func Sum(numbers ...int) int {
	defer fmt.Println("Sum function finished")

	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func safeDivision(a, b int) (int, error) {
	if b == 0 || a == 0 {
		return 0, &MathError{
			Operation: division,
			InputA:    a,
			InputB:    b,
			Message:   divisionError,
		}
	}

	return a / b, nil
}

func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	value, err := safeDivision(10, 2)
	fmt.Println(value, err)
}
