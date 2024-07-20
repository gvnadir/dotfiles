package main

import "fmt"

// The purpose of this example is to only show how to deal with pointers, but it's
// useless to use pointers for such a small data like a integer ("age") in this case
func main() {
	age := 32 // regular variable

	var agePointer *int // the "*" defines a type for a pointer
	agePointer = &age   // &age is the pointer of the "age" variable

	fmt.Println("Age:", *agePointer) // de-referencing a pointer

	// adultYears := getAdultYears(agePointer)
	// fmt.Println(adultYears)

	getAdultYears(agePointer)
	fmt.Println(age, &age)
}

// a temporary copy of the variable "age" will be passed here.
// the go garbage collector will remove that variable when the function completes its execution.
// func getAdultYears(age int) int {
// 	return age - 18
// }

// a pointer pointing to the same age stored value is passed here
// here we are returning a new value though
// func getAdultYears(age *int) int {
// 	return *age - 18
// }

// in this example we are overriding the "age" variable using its pointer
// without returning a new value
func getAdultYears(age *int) {
	fmt.Println(age)
	*age = *age - 18
}
