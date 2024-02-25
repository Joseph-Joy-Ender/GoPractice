package main

import "fmt"

func gasMileage() {
	fmt.Println("Enter miles driven: ")
	var miles int
	fmt.Scanln(&miles)

	fmt.Println("Enter gallons used for each tankful: ")
	var gallons int
	fmt.Scanln(&gallons)

	var totalMilesPerGallon float64 = 0

	var milesPerGallon = float64(miles / gallons)
	fmt.Printf("%.2f\n", milesPerGallon)

	totalMilesPerGallon += milesPerGallon

	var condition = true

	fmt.Println("Do you wish to continue? ")
	var answer string
	fmt.Scanln(&answer)

	if answer == "no" {
		condition = false
	}

	for condition {
		fmt.Println("Enter miles driven: ")
		fmt.Scanln(&miles)

		fmt.Println("Enter gallons used for each tankful: ")
		fmt.Scanln(&gallons)

		milesPerGallon = float64(miles / gallons)
		fmt.Println(milesPerGallon)

		totalMilesPerGallon += milesPerGallon

		fmt.Println("Do you wish to continue: ")
		fmt.Scanln(&answer)

		if answer == "no" {
			condition = false
		}
	}
	fmt.Printf("Total miles per gallon is %.4f ", totalMilesPerGallon)

}
