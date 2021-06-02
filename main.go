package main

import (
	"fmt"
	"go_scrapper/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err1 := dictionary.Add("first", "first word")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(dictionary)
	}
	definition, err2 := dictionary.Search("first")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}
	err3 := dictionary.Add("second", "second word")
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(dictionary)
	}

	err4 := dictionary.Update("second", "second update word")
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(dictionary)
	}

	err5 := dictionary.Delete("second")
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println(dictionary)
	}
}
