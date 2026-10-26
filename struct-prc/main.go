package main

import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}

type Account struct {
    HolderName string
    Balance    int
}

func main(){
	// result := Person{
	// 	Name: "rahul",
	// 	Age:123,
	// }

	// r1 := &result
	// r1.Age = 545
	// fmt.Printf("result %#v\n",result)
	// fmt.Println(result.Age)
	// fmt.Println(r1.Age)


	// accountDate := Account{
	// 	HolderName:"Rahul",
	// 	Balance: 40,
	// }

	// accCopy := &accountDate

	// accCopy.Balance = 5000

	// fmt.Println(accountDate.Balance)
	// fmt.Println(accCopy.Balance)

	// balance := 5000
	// deposit(&balance)
	// fmt.Println("this is the balnce", balance)

	//     account := Account{
    //     HolderName: "Rahul",
    //     Balance:    1000,
    // }

    // deposit1(&account, 500)

    // fmt.Println(account.Balance)

	accDetails := Account{
		HolderName: "abc",
		Balance: 3242,
	}

	acc2 := accDetails

	// acc2.updateDetails(5000)
	// accDetails.updateDetails(600)
	// fmt.Println("hdfsf",accDetails.Balance)

	// fmt.Println("this is ipdate", acc2.Balance)
	acc2.normal(500)
	fmt.Println("normal", acc2.Balance)
	accDetails.normal(2342)
	fmt.Println("sfsdfs", accDetails.Balance)
}

func(d *Account)updateDetails(bal int){
	d.Balance += bal
}

func (d Account)normal(bal int){
	d.Balance +=  bal
}
func deposit1(acc *Account, amount int) {
    // TODO: Increase the account balance by amount
	acc.Balance = acc.Balance + amount	
}
func deposit(balance *int){
	*balance = *balance + 500
}