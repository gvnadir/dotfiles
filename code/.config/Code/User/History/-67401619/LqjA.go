package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your account balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var deposit float64
		fmt.Scan(&deposit)
		accountBalance += deposit
		fmt.Println("Balance updated! New amount:", accountBalance)

	}

	fmt.Println("Your choice:", choice)
}
