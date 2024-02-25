package main

import "fmt"

func rightAngledTriangle() {
	fmt.Println("Enter length of triangle: ")
	var length int
	fmt.Scanln(&length)
	for i := 0; i <= length; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
