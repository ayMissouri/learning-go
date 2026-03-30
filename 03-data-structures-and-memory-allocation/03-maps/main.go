package main

import "fmt"

func main() {
	studentGrades := map[string]int{
		"John": 90,
		"Mark": 85,
		"Bob":  60,
	}
	fmt.Printf("%v\n", studentGrades)
	studentGrades["John"] = 95
	fmt.Printf("%v\n", studentGrades)

	// ok check, if its true then the key exists
	John, ok := studentGrades["John"]
	if ok {
		fmt.Printf("John's grade is %d\n", John)
	}

	// same line ok check
	if _, ok := studentGrades["Bob"]; ok {
		fmt.Printf("Bob's grade is %d\n", studentGrades["Bob"])
	}

	// delete a key
	delete(studentGrades, "Bob")
	fmt.Printf("%v\n", studentGrades)

	// use \:= or make to initialize
	var configs map[string]string
	// %T prints the type of the variable, in this case map[string]string
	fmt.Printf("%v %T\n", configs, configs)

	if configs == nil {
		fmt.Println("configs is nil")
	}
}
