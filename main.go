package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	length, _ := lenAndUpper("ilhoon")
	fmt.Println(multiply(2, 3))
	fmt.Println(length)
	repeatMe("aaa", "bbb", "ccc")
}
