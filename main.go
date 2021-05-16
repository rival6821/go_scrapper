package main

import "fmt"

func main() {
	names := [5]string{"a", "b"}
	names[2] = "11"
	names[3] = "22"
	names[4] = "33"

	nameSlice := []string{"vvv"}
	nameSlice = append(nameSlice, "bb")
	fmt.Println(names, nameSlice)
}
