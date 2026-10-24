package main

import "fmt"

type Account struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

type Bank struct{
	Accounts map[string]*Account
}

func main() {
	details := Account{AccountNumber: "123", HolderName: "abc", Balance: 1234.23}

	fmt.Println(details.AccountNumber)
	fmt.Println(details.Balance)

	bankAccounts :=  Bank{
		Accounts: make(map[string]*Account),
	}
	bankAccounts.Accounts[details.AccountNumber] = &details
	fmt.Println(bankAccounts.Accounts["123"].HolderName)

	nums:=500
	chageNum(nums)
	byref(&nums)

	fmt.Println(nums)
}

func byref(num *int) {
	*num = 500
}
func chageNum(num int){
	num =100
}

