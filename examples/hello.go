package main

import "fmt"

func main() {
	fmt.Println()
	printing()

}

func printing() {
	fmt.Println("Hello World!!")
	name := "Joy"
	fmt.Println("Hello my name is", name)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello" + " " + "World")
}
