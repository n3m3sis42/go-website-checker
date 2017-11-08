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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkSite(link, c)
	}

	for l := range c {
		// for {  <- above line does the exact same thing as this - creates an infinite loop - but is much clearer
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkSite(link, c)
		}(l) // we need to pass l into the function literal when we call it since l is declared outside the scope of the function literal
	}

}

func checkSite(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " might be down!")
		c <- link
		return
	}

	fmt.Println(link + " is up!")
	c <- link
}
