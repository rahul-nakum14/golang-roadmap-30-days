package main

import (
	"time"
	"errors"
	"fmt"
)

type TranscationType string

const (
	Deposit    TranscationType = "deposit"
	Withdrawal TranscationType = "withdrawal"
	Transfer   TranscationType = "transfer"
)


type Transaction struct {
	ID     string
	Amount float64
	Type   TranscationType
	Timestamp time.Time
}

type Account struct {
	AccountNumber           string
	HolderName              string
	Balance      float64
	Transactions []Transaction
}

type Bank struct {
	Accounts map[string]*Account
}

func main() {

	acc := &Account{
		AccountNumber: "123",
		HolderName:"Rahul",
		Balance: 1000,
	}

	if err := acc.Deposit(500); err != nil{
		fmt.Println("Deposit error", err)
		return 
	}

	if err := acc.Withdrawal(200); err != nil {
		fmt.Println("Withdrawal error:", err)
		return
	}

	acc.Getstatement()
}

func(d *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("Deposit must be greater than 0")
	}
	d.Balance += amount
	d.Transactions = append(d.Transactions, Transaction{ID:"1", Amount: amount , Type: Deposit, Timestamp: time.Now()})
	return nil
}

func(w *Account) Withdrawal(amount float64) error {
	if amount > w.Balance {
		return  fmt.Errorf("insufficient funds: balance is %.2f, attempted %.2f", w.Balance, amount)
	}
	w.Balance -= amount
	w.Transactions = append(w.Transactions, Transaction{ID:"2", Amount: amount , Type: Withdrawal, Timestamp: time.Now()})
	return nil
}

func (s *Account)Getstatement(){
	// fmt.Println("",s.AccountNumber,s.Balance,s.HolderName,s.Transactions)
	fmt.Printf(
		"Account: %s | Holder: %s | Balance: %.2f\n",
		s.AccountNumber,
		s.HolderName,
		s.Balance,
	)

		fmt.Println("Transactions:")

	// for _, transaction := range s.Transactions {
	// 	fmt.Printf(
	// 		"ID: %s | Type: %s | Amount: %.2f | Time: %s\n",
	// 		transaction.ID,
	// 		transaction.Type,
	// 		transaction.Amount,
	// 		transaction.Timestamp.Format(time.RFC3339),
	// 	)
	// }
}


func NewBank() *Bank{
	return &Bank{
		Accounts: make(map[string]*Account),
	}
}
