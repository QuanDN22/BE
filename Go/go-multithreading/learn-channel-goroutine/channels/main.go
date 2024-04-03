package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://github.com",
		"http://amazon.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for {
		go checkLink(<-ch, ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(err)
		fmt.Println(link, "might be down")
		ch <- link
		return
	}

	fmt.Println(link, "is up")
	ch <- link
}
