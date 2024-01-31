package channel

import (
	"fmt"
	"net/http"
	"time"
)

func IsSexy(person string, channel chan string) {
	//time.Sleep(time.Second * 5)
	channel <- person + " is Sexy!!"
}

func SexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

type requestResult struct {
	url    string
	status string
}

func UrlChecker() {
	url := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	results := map[string]string{}
	channel := make(chan requestResult)
	for _, url := range url {
		go hitURL(url, channel)
	}

	for index, _ := range url {
		result := <-channel
		results[result.url] = result.status
		_ = index
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, channel chan<- requestResult) {
	//fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	channel <- requestResult{url: url, status: status}
}
