package main

import (
	"fmt"
	"os"
)

type userLoginDetails struct {
	userId   string
	password string
}

var balance int

func selectAnOption() {
	fmt.Println("Please select an option from the menu below:\n")
	fmt.Println("w -> Withdraw money\n")
	fmt.Println("d -> Deposit money\n")
	fmt.Println("r -> Request balance\n")
	fmt.Println("e -> Exit the program\n")

	var operation string
	fmt.Scanln(&operation)

	if operation == "w" {
		WithdrawMoney()
	} else if operation == "d" {
		DepositMoney()
	} else if operation == "r" {
		RequestBalance()
	} else if operation == "e" {
		os.Exit(1)
	}

}
func WithdrawMoney() {
	if balance == 0 {
		fmt.Println("In-sufficient funds in your account")
		selectAnOption()
	} else {
		fmt.Print("Enter the amount to withdraw: ")
		fmt.Scan(&balance)

		fmt.Println(" withdraw successfully.", "\n", "Your current balance is : ", balance)
		selectAnOption()
	}
}
func DepositMoney() {
	fmt.Print("Enter Amount you want to deposit : ")
	fmt.Scan(&balance)

	fmt.Println("Amount deposit successfully. Your current balance is : ", balance)
	selectAnOption()
}

func RequestBalance() {
	fmt.Println("Current Balance is : ", balance)
	selectAnOption()
}

func main() {

	fmt.Println("Hi! Welcome to Mr.Vijay ATM Machine!\n")
	fmt.Println("Please select an option from the menu below")

	var option string

	fmt.Println("l  -> Login")
	fmt.Println("c -> Create New Account")
	fmt.Println("q -> Quit")
	fmt.Scan(&option)

	var id string
	var password string

	if option == "l" {
		fmt.Print("Please enter userid : ")
		fmt.Scan(&id)

		fmt.Print("Please enter your password : ")
		fmt.Scan(&password)

		s := userLoginDetails{userId: id, password: password}

		fmt.Println("Your userId is", s.userId, "and password is", password)

		selectAnOption()

	} else if option == "c" {
		fmt.Print("Please enter userid : ")
		fmt.Scan(&id)

		fmt.Print("Please enter your password : ")
		fmt.Scan(&password)

		s := userLoginDetails{userId: id, password: password}

		fmt.Println("Your Account has been created with userId ", s.userId)
	} else if option == "q" {
		os.Exit(1)
	}
}
