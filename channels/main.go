package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://twitter.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, l := range links {
		go statusChecker(l, c)

	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go statusChecker(link, c)
		}(l)
	}

}

func statusChecker(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link

}
