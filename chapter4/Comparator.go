package main

import "fmt"

func compare(firstNumber int, secondNumber int) {
	if firstNumber == secondNumber {
		fmt.Println(0)
	}
	if firstNumber > secondNumber {
		fmt.Println(1)
	}
	if secondNumber > firstNumber {
		fmt.Println(-1)
	}
}
