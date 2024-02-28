package main

import "fmt"

type creditLimitCalculator struct {
	accountNumber    int
	beginningBalance int
	totalItems       int
	totalOfCredits   int
	creditLimit      int
}

func (limit *creditLimitCalculator) getAccountNumber() int {
	return limit.accountNumber
}

func (limit *creditLimitCalculator) setAccountNumber(accountNumber int) {
	limit.accountNumber = accountNumber
}

func (limit *creditLimitCalculator) getBeginningBalance() int {
	return limit.beginningBalance
}

func (limit *creditLimitCalculator) setBeginningBalance(beginningBalance int) {
	limit.beginningBalance = beginningBalance
}

func (limit *creditLimitCalculator) getTotalItems() int {
	return limit.totalItems
}

func (limit *creditLimitCalculator) setTotalItems(totalItems int) {
	limit.totalItems = totalItems
}

func (limit *creditLimitCalculator) getTotalOfCredits() int {
	return limit.totalOfCredits
}

func (limit *creditLimitCalculator) setTotalOfCredits(totalOfCredits int) {
	limit.totalOfCredits = totalOfCredits
}

func (limit *creditLimitCalculator) getCreditLimit() int {
	return limit.creditLimit
}

func (limit *creditLimitCalculator) setCreditLimit(creditLimit int) {
	limit.creditLimit = creditLimit
}

func calculateNewBalance() int {
	var limit = &creditLimitCalculator{}
	newBalance := limit.getBeginningBalance() + limit.getTotalItems() - limit.getTotalOfCredits()
	displayMessage(newBalance)
	return newBalance
}

func displayMessage(newBalance int) {
	var limit *creditLimitCalculator
	limit = &creditLimitCalculator{}
	if newBalance > limit.getCreditLimit() {
		fmt.Println("Credit limit exceeded")
	}

}
