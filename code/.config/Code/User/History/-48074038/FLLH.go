package main

import "fmt"

// input: revenue, expenses, rate
// calculate earning before tax (EBT) and earning after tax (profit)
// calculate ratio (EBT / profit)
// output: EBT, profit and ratio

const rate = 20.0

func main() {
	revenue := readAndStore("Revenue: ")
	expenses := readAndStore("Expenses: ")

	var ebt, profit, ratio = calculateProfit(revenue, expenses)
	fmt.Println("EBT:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}

func readAndStore(text string) float64 {
	var out float64
	fmt.Print(text)
	fmt.Scan(out)
	return out
}

func calculateProfit(revenue, expenses float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	taxAmount := ebt * (rate / 100)
	profit = ebt - taxAmount
	ratio = ebt / profit
	return ebt, profit, ratio
}
