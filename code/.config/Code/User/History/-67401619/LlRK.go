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

		switch choice {
		case 1:
			fmt.Println("Your account balance is:", accountBalance)
			continue
		case 2:
			fmt.Print("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += deposit
			fmt.Println("Balance updated! New amount:", accountBalance)
			continue
		case 3:
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
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
