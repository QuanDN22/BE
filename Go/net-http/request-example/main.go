package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// START OMIT
	// Create a new request to the specified URL
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	fmt.Println(req.Header.Get("Accept"))
	// END OMIT

	// START OMIT
	// Create a new request to the specified URL
	resp, err := http.Get("http://facebook.com")
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	fmt.Println(resp.StatusCode)

	resp, err = http.PostForm("http://example.com/form",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	fmt.Println(resp.StatusCode)

	// END OMIT

	// START OMIT
	// Create a new request to the specified URL

	resp, err = http.Get("http://204.106.240.53")
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	fmt.Println(resp.StatusCode)

	

	resp, err = http.Get("http://203.126.118.38")
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	fmt.Println(resp.StatusCode)

	// END OMIT

}
