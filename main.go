package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Input revenue  ")
	fmt.Scan(&revenue)

	fmt.Print("Input expenses ")
	fmt.Scan(&expenses)

	fmt.Print("Input tax rate ")
	fmt.Scan(&tax_rate)

	earning_before_tax := expenses / tax_rate
	profit := revenue / tax_rate
	ratio := expenses

	fmt.Println("EBT    %.1f", earning_before_tax)
	fmt.Println("Profit %.1f", profit)
	fmt.Println("Ratio  %.1f", ratio)
}
