package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	// try to remove the balance.txt file in order to get the error message and the fallback value
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data) // string() here is used for convert byte[] to string
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return value, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance) // fmt.Sprint() is used for convert float64 to string
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------------------")

		// the line below will stop the execution and throw an error
		// panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
