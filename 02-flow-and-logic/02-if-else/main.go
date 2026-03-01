package main

import "fmt"

func main() {
	temp := 25

	if temp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("less than 30")
	}

	userAccess := map[string]bool{
		"admin": true,
		"user":  false,
	}

	if hasAccess, ok := userAccess["admin"]; ok && hasAccess {
		fmt.Println("admin access")
	} else {
		fmt.Println("no access")
	}
}
