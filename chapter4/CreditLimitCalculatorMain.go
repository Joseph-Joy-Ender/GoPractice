package main

import "fmt"

func creditMain() {
	credit := new(creditLimitCalculator)

	fmt.Println("Enter your account number: ")
	var accountNumber int
	fmt.Scanln(&accountNumber)
	credit.setAccountNumber(accountNumber)

	fmt.Println("Enter beginning balance: ")
	var beginningBalance int
	fmt.Scanln(&beginningBalance)
	credit.setBeginningBalance(beginningBalance)

	fmt.Println("Enter total of all items: ")
	var totalItems int
	fmt.Scanln(&totalItems)
	credit.setTotalItems(totalItems)

	fmt.Println("Enter total of credits: ")
	var totalOfCredits int
	fmt.Scanln(&totalOfCredits)
	credit.setTotalOfCredits(totalOfCredits)

	fmt.Println("Enter credit limit: ")
	var creditLimit int
	fmt.Scanln(&creditLimit)
	credit.setCreditLimit(creditLimit)

	fmt.Println("Account number is ", credit.getAccountNumber())
	fmt.Println("Your balance at the beginning of the month is ", credit.getBeginningBalance())
	fmt.Println("Total of all items charged by customers is ", credit.getTotalItems())
	fmt.Println("Total of all credits applied to account is ", credit.getTotalOfCredits())
	fmt.Println("The allowed credit limit is ", credit.getCreditLimit())
	fmt.Println("The new balance is ", calculateNewBalance())

}
