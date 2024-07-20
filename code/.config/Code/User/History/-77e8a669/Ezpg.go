package main

import "fmt"

func main() {
	age := 32 // regular variable

	fmt.Println("Age:", age)

	// a temporary copy of the variable "age" will be passed here.
	// the go garbage collector will remove that variable when the function completes its execution.
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
