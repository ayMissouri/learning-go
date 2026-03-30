package main

import "fmt"

func modifyValue(val int) {
	val = val * 10
	fmt.Printf("modifyValue: %v\n", val)
}

// * to specify that the pointer is a pointer
func modifyPointer(val *int) {
	if val == nil {
		fmt.Printf("val is nil\n")
		return
	}

	// dereference the pointer
	// go to the address, multiply by 10, and assign it to the original variable
	*val = *val * 10
	fmt.Printf("modifyPointer: %v\n", *val)
}

func main() {
	// a pointer is a variable that holds the address of another variable
	//age \:= 10
	//agePtr \:= &age
	//fmt.Printf("AgePtr: %d\n", agePtr)
	//fmt.Printf("Age address: %d\n", &age)

	num := 10
	modifyValue(num)
	fmt.Printf("num: %d\n", num)

	// &num to pass the address of num
	modifyPointer(&num)
	fmt.Printf("num: %d\n", num)
	
	grade := 50
	gradePtr := &grade
	fmt.Printf("grade: %v\n", gradePtr)
	fmt.Printf("gradePtr: %v\n", &gradePtr)
}
