package main

import (
	"fmt"
	"go_scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("ilhoon")
	fmt.Println(account)
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(110)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
