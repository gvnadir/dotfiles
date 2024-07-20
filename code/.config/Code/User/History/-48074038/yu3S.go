package main

import "fmt"

// input: revenue, expenses, rate
// calculate earning before tax (EBT) and earning after tax (profit)
// calculate ratio (EBT / profit)
// output: EBT, profit and ratio

func main() {
	revenue := readAndStore("Revenue: ")
	expenses := readAndStore("Expenses: ")
	rate := readAndStore("Tax Rate: ")

	var ebt, profit, ratio = calculateProfit(revenue, expenses, rate)
	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}

func readAndStore(text string) (out float64) {
	fmt.Print(text)
	fmt.Scan(&out)
	return out
}

func calculateProfit(revenue, expenses, rate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	taxAmount := ebt * (rate / 100)
	profit = ebt - taxAmount
	ratio = ebt / profit
	return ebt, profit, ratio
}
