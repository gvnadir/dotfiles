package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int // the "*" defines a type for a pointer
	agePointer = &age   // &age is the pointer of the "age" variable

	fmt.Println("Age:", age)

	// a temporary copy of the variable "age" will be passed here.
	// the go garbage collector will remove that variable when the function completes its execution.
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
