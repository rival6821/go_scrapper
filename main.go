package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kk", "bb"}
	tt := person{name: "ilhon", age: 22, favFood: favFood}
	fmt.Println(tt)
}
