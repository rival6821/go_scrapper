package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("REQUEST ERROR")

func main() {
	urls := []string{
		"https://naver.com",
		"https://ilhoon.kr",
		"https://youtube.com",
	}
	for _, url := range urls {
		err := hitURL(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	} else if resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
