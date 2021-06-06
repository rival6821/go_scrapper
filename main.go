package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)

	c := make(chan requestResult)

	urls := []string{
		"https://naver.com",
		"https://ilhoon.kr",
		"https://youtube.com",
		"https://nomadcoders.co",
		"https://server.ilhoon.kr",
		"https://land.naver.com",
		"https://comic.naver.com/index.nhn",
		"https://www.clien.net/service/board/park",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	// fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
