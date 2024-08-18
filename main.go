package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	tax_rate := getUserInput("Tax rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)

	fmt.Printf("EBT%.1f\n", ebt)
	fmt.Printf("Profit%.1f\n", profit)
	fmt.Printf("Ratio%.1f\n", ratio)
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}
