package main

import (
	"fmt"
	"math"
)

func interest() {
	fmt.Print("Enter principal: ")
	var principal float64
	fmt.Scanln(principal)

	fmt.Print("Enter rate: ")
	var rate float64

	fmt.Printf("%s%23s\n", "Year", "Amount on deposit")

	for year := 1; year <= 10; year++ {

		amount := principal * math.Pow(1.0+rate, float64(year))

		fmt.Printf("%4d%20.2f\n", year, amount)

	}
}
