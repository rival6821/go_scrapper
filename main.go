package main

import "fmt"

func main() {
	tt := map[string]string{"aa": "bb", "age": "12"}
	for key, value := range tt {
		fmt.Println(key, value)

	}
}
