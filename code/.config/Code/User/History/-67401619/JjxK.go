package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("You choice: ")
		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is:", accountBalance)
			continue
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			accountBalance += deposit
			fmt.Println("Balance updated! New amount:", accountBalance)
			continue
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdraw > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdraw
			fmt.Println("Balance updated! New amount:", accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for choosing our bank!")
}
