package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	length, uppercase := lenAndUpper("ilhoon")
	fmt.Println(multiply(2, 3))
	fmt.Println(length, uppercase)
	repeatMe("aaa", "bbb", "ccc")
}
