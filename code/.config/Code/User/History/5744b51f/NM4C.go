package main

import "fmt"

// input: revenue, expenses, rate
// calculate earning before tax (EBT) and earning after tax (profit)
// calculate ratio (EBT / profit)
// output: EBT, profit and ratio

func main() {
	var revenue int
	var expenses int
	rate := 20.0

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	ebt := revenue - expenses
	taxAmount := revenue * int(rate)
	profit := ebt - taxAmount
	ratio := ebt / profit

	fmt.Print("EBT: ")
	fmt.Print(&ebt)
	fmt.Print("profit: ")
	fmt.Print(&profit)
	fmt.Print("ratio: ")
	fmt.Print(&ratio)

}
