package main

import "fmt"

func main() {
	fmt.Println("Enter your account number: ")
	var accountNumber string
	fmt.Scanln(&accountNumber)
	fmt.Println("Enter your balance at the beginning of the month: ")
	var accountBalance int
	fmt.Scanln(&accountBalance)
	fmt.Println("Enter the total of all items charged by the customer this month: ")
	var totalCustomerCharge int
	fmt.Scanln(&totalCustomerCharge)
	fmt.Println("Enter the total of all credits applied to the customer account this month: ")
	var customerCredit int
	fmt.Scanln(&customerCredit)
	fmt.Println("Enter the allowed credit limit: ")
	var creditLimit int
	fmt.Scanln(&creditLimit)
	newBalance := accountBalance + totalCustomerCharge - customerCredit
	fmt.Println("The new balance is ", newBalance)
	if newBalance > customerCredit {
		fmt.Println("Still Eligible :)! ")
	} else {
		fmt.Println("Credit limit exceeded :(!!")
	}

}
