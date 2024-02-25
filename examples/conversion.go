package main

import "fmt"

func main() {
	conversion()

}

func conversion() {
	fmt.Println("Enter fahrenheit: ")
	var fah float64
	fmt.Scanf("%f", &fah)

	formula := fah - 32*5/9

	fmt.Println(formula)

}
