package main

import "fmt"

func shape1() {
	for i := 1; i <= 5; i++ {
		for space := i; space < 6; space++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}
}

func shape2() {
	for i := 0; i < 6; i++ {
		fmt.Print("* ")

	}
	fmt.Println()
}

func shape3() {
	for i := 5; i > 0; i-- {
		for spaces := i; spaces < 6; spaces++ {
			fmt.Print(" ")
		}
		for j := i; j > 0; j-- {
			fmt.Print("* ")

		}
		fmt.Println()
	}
}
