package main

import "fmt"

func SalesCommissionMain() {
	sales := SalesCommission{}

	var totalOfItemSold float64 = 0

	var condition bool = true

	fmt.Println("Enter price for item sold: ")
	var priceOfItems float64
	fmt.Scanln(&priceOfItems)

	totalOfItemSold += priceOfItems

	fmt.Println("Do you wish to continue: ")
	var answer string
	fmt.Scanln(&answer)

	if answer == "no" {
		condition = false
	}

	for condition {
		fmt.Println("Enter price of item sold: ")
		fmt.Scanln(&priceOfItems)

		totalOfItemSold += priceOfItems

		fmt.Println("Do you wish to continue: ")
		fmt.Scanln(&answer)

		if answer == "no" {
			condition = false
		}
	}

	sales.setPriceForItemSold(totalOfItemSold)
	fmt.Println("Total of item sold is ", sales.getPriceForItemSold())
	fmt.Println("Salary is ", getSalary())
	fmt.Println("Percentage is ", getPercentage())
	fmt.Println("Calculated commission is ", calculateCommission())
	fmt.Println("Total wages is ", calculateTotalWage())
}
