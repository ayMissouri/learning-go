package main

import (
	"fmt"
	"time"
)

func main() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("Beginning of the week")
	case "tuesday", "wednesday", "thursday":
		fmt.Println("It's a weekday")
	case "friday":
		fmt.Println("Last day of the week")
	default:
		fmt.Println("It's the weekend!")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("Int", v)
		case string:
			fmt.Println("String", v)
		case bool:
			fmt.Println("Boolean", v)
		default:
			fmt.Println("Unknown type")
		}
	}

	checkType(10)
	checkType("hello")
	checkType(true)
	checkType(10.5)
}
