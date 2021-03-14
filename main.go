package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fetch(url string) {
	defer wg.Done()
	start_time := time.Now()
	resp, err := http.Get(url)
	end_time := time.Since(start_time)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Url: %s | Length: %d | Processing time: %s\n", url, len(body), end_time)
}

func fetchWithChannels(urls []string, ch chan string) {
	for _, url := range urls {
		start_time := time.Now()
		resp, err := http.Get(url)
		end_time := time.Since(start_time)
		if err != nil {
			fmt.Println(err)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		res := fmt.Sprintf("Url: %s | Length: %d | Processing time: %s\n", url, len(body), end_time)
		ch <- res
	}
	close(ch)
}

func main() {
	urls := []string{"https://imadelmahrad.com", "https://space.imadelmahrad.com", "https://chillingbook.fr", "https://github.com"}
	size := len(urls)
	wg.Add(size)
	for _, url := range urls {
		go fetch(url)
	}
	wg.Wait()

	fmt.Println("--------------------------------")

	ch := make(chan string, size)
	go fetchWithChannels(urls, ch)

	ok := true

	for ok {
		select {
		case s, open := <-ch:
			{
				if !open {
					ok = false
				}
				fmt.Println(s)
			}
		}
	}
}
