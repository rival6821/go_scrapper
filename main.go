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
	account.Withdraw(10)
	fmt.Println(account.Balance())
}
