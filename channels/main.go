package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://kodega2016.com",
	}
	ch := make(chan string)
	for _, link := range links {
		go checkWeb(link, ch)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// }

	for l := range ch {
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkWeb(link, ch)
		}(l)
	}
}

func checkWeb(link string, ch chan string) {
	_, err := http.Get(link)
	ch <- link
	if err != nil {
		// ch <- fmt.Sprintf("%s might be down", link)
		fmt.Println(link, "might be down")
		// return
	}
	fmt.Println(link, "is up")
	// ch <- fmt.Sprintf("%s is up", link)
}
