package main

import (
	"fmt"
)

func canIDrink(age int) bool {

	if koreaAge := age + 2; koreaAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink((16)))
}
