package main

import "fmt"

func checkerBoard() {
	for i := 1; i < 9; i++ {
		if i%2 == 0 {
			fmt.Print(" ")
		}
		for j := 0; j <= 8; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}
