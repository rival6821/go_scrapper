package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	total := superAdd(1, 2, 3, 4, 5)
	fmt.Println(total)
}
