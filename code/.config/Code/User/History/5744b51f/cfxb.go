package main

import "fmt"

// input: revenue, expenses, rate
// calculate earning before tax (EBT) and earning after tax (profit)
// calculate ratio (EBT / profit)
// output: EBT, profit and ratio

func main() {
	var revenue, expenses float64
	rate := 20.0

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	ebt := revenue - expenses
	taxAmount := revenue * (rate / 100)
	profit := ebt - taxAmount
	ratio := ebt / profit

	fmt.Print("EBT: ")
	fmt.Println(ebt)

	fmt.Print("profit: ")
	fmt.Println(profit)

	fmt.Print("ratio: ")
	fmt.Println(ratio)

}
