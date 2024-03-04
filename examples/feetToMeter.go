package main

import "fmt"

func feetToMeter() {
	fmt.Println("Enter feet")
	var feet float64

	fmt.Scanf("%f", &feet)
	formula := feet * 0.3048

	fmt.Println(formula)
}
