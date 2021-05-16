package main

import (
	"fmt"
	"go_scrapper/accounts"
)

func main() {
	account := accounts.NewAccount("ilhoon")
	fmt.Println(account)
}
